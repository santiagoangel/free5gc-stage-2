#include "upf_config.h"

#include "upf_context.h"
#include "utlt_yaml.h"
#include "gtp_link.h"

static int SetProtocolIter(YamlIter *protoList, YamlIter *protoIter);
static void DeleteYamlDocument();

static Status AddGtpv1Endpoint(const char *host); // host: hostname or ip address
static Status AddPfcpEndpoint(const char *host); // host: hostname or ip address
static Status AddGtpv1EndpointWithName(const char *host, const char *ifname);

static yaml_document_t *document = NULL;

Status UpfLoadConfigFile(const char *configFilePath) {
    FILE *file;
    yaml_parser_t parser;

    file = fopen(configFilePath, "rb");
    UTLT_Assert(file, return STATUS_ERROR, "Fail to open yaml file");

    UTLT_Assert(yaml_parser_initialize(&parser), return STATUS_ERROR, "Fail to initialize parser");
    yaml_parser_set_input_file(&parser, file);

    document = UTLT_Calloc(1, sizeof(yaml_document_t));
    
    if (!yaml_parser_load(&parser, document)) {
        UTLT_Free(document);
        yaml_parser_delete(&parser);
        UTLT_Assert(!fclose(file), return STATUS_ERROR, "Fail to close yaml file");
    }

    return STATUS_OK;
}

Status UpfConfigParse() {
    UTLT_Assert(document, return STATUS_ERROR, "Config not loaded");

    YamlIter rootIter;

    YamlIterInit(&rootIter, document);
    while (YamlIterNext(&rootIter)) {
        const char *rootKey = YamlIterGet(&rootIter, GET_KEY);
        UTLT_Assert(rootKey, return STATUS_ERROR, "The rootKey is NULL");
        
        if (!strcmp(rootKey, "configuration")) {
            YamlIter upfIter;
            YamlIterChild(&rootIter, &upfIter);
            while (YamlIterNext(&upfIter)) {
                const char *upfKey = YamlIterGet(&upfIter, GET_KEY);
                UTLT_Assert(upfKey, return STATUS_ERROR, "The rootKey is NULL");

                if (!strcmp(upfKey, "debugLevel")) {
                    const char *logLevel = YamlIterGet(&upfIter, GET_VALUE);
                    
                    UTLT_Assert(UTLT_SetLogLevel(logLevel) == STATUS_OK,
                                return STATUS_ERROR, "");

                } else if (!strcmp(upfKey, "gtpu")) {
                    YamlIter gtpuList, gtpuIter;
                    YamlIterChild(&upfIter, &gtpuList);

                    do {
                        const char *host;
                        int port;
                        const char *ifname = NULL;

                        if (SetProtocolIter(&gtpuList, &gtpuIter)) {
                            break;
                        }

                        while (YamlIterNext(&gtpuIter)) {
                            const char *gtpuKey = YamlIterGet(&gtpuIter, GET_KEY);
                            UTLT_Assert(gtpuKey, return STATUS_ERROR, "The gtpuKey is NULL");

                            if (!strcmp(gtpuKey, "addr") || !strcmp(gtpuKey, "name")) {
                                host = YamlIterGet(&gtpuIter, GET_VALUE);
                            } else if (!strcmp(gtpuKey, "family")) {
                                // TODO: support IPv6
                            } else if (!strcmp(gtpuKey, "port")) {
                                const char *v = YamlIterGet(&gtpuIter, GET_KEY);
                                if (v) {
                                    port = atoi(v);
                                    Self()->gtpv1Port = port;
                                }
                            } else if (!strcmp(gtpuKey, "ifname")) {
                                ifname = (char *)YamlIterGet(&gtpuIter, GET_VALUE);
                            } else {
                                UTLT_Warning("Unknown key \"%s\" of gtpu", gtpuKey);
                            }
                        }

                        if (host) {
                            if (ifname)
                                AddGtpv1EndpointWithName(host, ifname);
                            else
                                AddGtpv1Endpoint(host);
                        }
                        
                    } while (YamlIterType(&gtpuList) == YAML_SEQUENCE_NODE);
                    
                } else if (!strcmp(upfKey, "pfcp")) {
                    YamlIter pfcpList, pfcpIter;
                    YamlIterChild(&upfIter, &pfcpList);

                    do {
                        const char *host;
                        
                        if (SetProtocolIter(&pfcpList, &pfcpIter))
                            break;
                        
                        while (YamlIterNext(&pfcpIter)) {
                            const char *pfcpKey = YamlIterGet(&pfcpIter, GET_KEY);
                            UTLT_Assert(pfcpKey, return STATUS_ERROR, "The pfcpKey is NULL");

                            if (!strcmp(pfcpKey, "addr") || !strcmp(pfcpKey, "name")) {
                                host = YamlIterGet(&pfcpIter, GET_VALUE);
                            } else {
                                UTLT_Warning("Unknown key \"%s\" of pfcp", pfcpKey);
                            }
                        }

                        if (host)
                            AddPfcpEndpoint(host);

                    } while (YamlIterType(&pfcpList) == YAML_SEQUENCE_NODE);
                    
                } else if (!strcmp(upfKey, "apn_list")) {
                    YamlIter apnList, apnIter;
                    YamlIterChild(&upfIter, &apnList);

                    do {
                        const char *apn = NULL;
                        const char *ipStr = NULL;
                        const char * mask = NULL;

                        if (SetProtocolIter(&apnList, &apnIter))
                            break;

                        while (YamlIterNext(&apnIter)) {
                            const char *apnKey = YamlIterGet(&apnIter, GET_KEY);
                            UTLT_Assert(apnKey, return STATUS_ERROR, "The apnKey is NULL");

                            if (!strcmp(apnKey, "cidr")) {
                                char *val = (char *)YamlIterGet(&apnIter, GET_VALUE);
                                
                                if (val) {
                                    ipStr = (const char *)strsep(&val, "/");
                                    if (ipStr)
                                        mask = (const char *)val;
                                }
                            } else if (!strcmp(apnKey, "apn")) {
                                apn = (char *)YamlIterGet(&apnIter, GET_VALUE);
                            } else {
                                UTLT_Warning("Unknown key \"%s\" of apn_list", apnKey);
                            }
                        }

                        if (apn && ipStr && mask) {
                            UTLT_Assert(!UpfApnAdd(apn, ipStr, mask), return STATUS_ERROR, "");
                        }
                    } while (YamlIterType(&apnList) == YAML_SEQUENCE_NODE);
                } else
                    UTLT_Warning("Unknown key \"%s\" of configuration", upfKey);
            }
        }
    }

    DeleteYamlDocument();
    
    return STATUS_OK;
}

static int SetProtocolIter(YamlIter *protoList, YamlIter *protoIter) {
    if (YamlIterType(protoList) == YAML_SCALAR_NODE) {
        return 1;
    } else if (YamlIterType(protoList) == YAML_SEQUENCE_NODE) {
        if (!YamlIterNext(protoList))
            return 1;
        YamlIterChild(protoList, protoIter);
    } else if (YamlIterType(protoList) == YAML_MAPPING_NODE) {
        memcpy(protoIter, protoList, sizeof(YamlIter));
    } else {
        UTLT_Assert(0, return 0, "Unknown node type");
        return 1;
    }

    return 0;
}

static void DeleteYamlDocument() {
    yaml_document_delete(document);
    UTLT_Free(document);
}

static Status AddGtpv1Endpoint(const char *host) {
    char ifname[MAX_IFNAME_STRLEN];
    sprintf(ifname, "%s%d", Self()->gtpDevNamePrefix, Self()->gtpv1DevSN++);

    return AddGtpv1EndpointWithName(host, ifname);
}

static Status AddGtpv1EndpointWithName(const char *host, const char *ifname) {
    UTLT_Assert(host, return STATUS_ERROR, "");

    int result;
    char ip[INET6_ADDRSTRLEN];

    result = GetAddrFromHost(ip, host, INET6_ADDRSTRLEN);
    UTLT_Assert(result == STATUS_OK, return STATUS_ERROR, "");

    Gtpv1TunDevNode *node = Gtpv1DevListAdd(&Self()->gtpv1DevList, ip, ifname);
    UTLT_Assert(node, return STATUS_ERROR, "");

    return STATUS_OK;
}

static Status AddPfcpEndpoint(const char *host) {
    UTLT_Assert(host, return STATUS_ERROR, "");

    int result;
    char ip[INET6_ADDRSTRLEN];

    result = GetAddrFromHost(ip, host, INET6_ADDRSTRLEN);
    UTLT_Assert(result == STATUS_OK, return STATUS_ERROR, "");

    SockNode *node = SockNodeListAdd(&Self()->pfcpIPList, ip);
    UTLT_Assert(node, return STATUS_ERROR, "");

    return STATUS_OK;
}
