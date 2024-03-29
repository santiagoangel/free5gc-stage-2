#include "up_gtp_path.h"

#include "utlt_debug.h"
#include "utlt_buff.h"
#include "utlt_network.h"
#include "gtp_header.h"
#include "gtp_path.h"
#include "upf_context.h"

Status GTPv1ServerInit() {
    Status status;

    status = GtpDevListCreate(Self()->epfd, AF_INET, &Self()->gtpv1DevList, GtpHandler, NULL);
    UTLT_Assert(status == STATUS_OK, return STATUS_ERROR, "GtpLinkListCreate fail");

    return STATUS_OK;
}

Status GTPv1ServerTerminate() {
    Status status = STATUS_OK;

    UTLT_Assert(GtpTunListFree(Self()->epfd, &Self()->gtpv1DevList) == STATUS_OK,
                status |= STATUS_ERROR, "GTPv1 tunnel list free fail");

    return status;
}

Status GtpHandler(Sock *sock, void *data) {
    UTLT_Assert(sock, return STATUS_ERROR, "GTP socket not found");
    Status status = STATUS_ERROR;

    Bufblk *pktbuf = BufblkAlloc(1, MAX_OF_GTPV1_PACKET_SIZE);
    int readNum = GtpRecv(sock, pktbuf);
    UTLT_Assert(readNum >= 0, goto FREEBUFBLK, "GTP receive fail");

    // TODO: Need to handle buffering and reject, including GTP and general packet
    // Not only GTP packet
    Gtpv1Header *gtpHdr = pktbuf->buf;
    UTLT_Assert(gtpHdr->version == 1, goto FREEBUFBLK,
                "Only handle the GTP version 1 in user plane");

    switch (gtpHdr->type) {
        case GTPV1_ECHO_REQUEST :
            status = GtpHandleEchoRequest(sock, gtpHdr);
            break;
        case GTPV1_ECHO_RESPONSE :
            status = GtpHandleEchoResponse(gtpHdr);
            break;
        case GTPV1_ERROR_INDICATION :

            break;
        case GTPV1_END_MARK :
            status = GtpHandleEndMark(sock, gtpHdr);
            break;
        case GTPV1_T_PDU :
            status = GtpHandleTPDU(sock, gtpHdr);
            break;
        default : 
            UTLT_Warning("This type[%d] of GTPv1 header does not implement yet", gtpHdr->type);
    }

FREEBUFBLK:
    UTLT_Assert(BufblkFree(pktbuf) == STATUS_OK, , "Bufblk free fail");

    return status;
}

Status GtpHandleEchoRequest(Sock *sock, void *data) {
    UTLT_Assert(data, return STATUS_ERROR, "GTP data is NULL");

    Gtpv1Header *gtpHdr = data;
    UTLT_Assert(gtpHdr->type == GTPV1_ECHO_REQUEST, return STATUS_ERROR, 
                "The type of GTP data is not 'Echo Request'");

    Status status = STATUS_OK;

    // Build the Echo Response packet
    Gtpv1Header gtpRespHrd = {
        .version = 1,
        .PT = 1,
        .type = GTPV1_ECHO_RESPONSE,
        .seqFlag = gtpHdr->seqFlag,
        .pn = gtpHdr->pn,
    };

    Bufblk *optPkt = BufblkAlloc(1, 0x40);
    if (gtpRespHrd.seqFlag | gtpRespHrd.pn) {
        Gtp1OptHeader *opthrd = (void *)((uint8_t *) data + GTPV1_HEADER_LEN);
        Gtp1OptHeader gtpOptHrd = {
            ._seqNum = gtpRespHrd.seqFlag ? htons(ntohs(opthrd->_seqNum) + 1) : 0,
            .nPdnNum = gtpRespHrd.pn ? opthrd->nPdnNum : 0,
        };
        BufblkBytes(optPkt, (void *) &gtpOptHrd, sizeof(gtpOptHrd));
    }

    /* Recover IE */
    uint8_t recoverType = 14, recoverCnt = 0;
    BufblkBytes(optPkt, (void *) &recoverType, 1);
    BufblkBytes(optPkt, (void *) &recoverCnt, 1);

    gtpRespHrd._length = htons(optPkt->len);

    Bufblk *pkt = BufblkAlloc(1, 0x40);
    BufblkBytes(pkt, (void *) &gtpRespHrd, GTPV1_HEADER_LEN);
    BufblkBuf(pkt, optPkt);

    BufblkFree(optPkt);

    UTLT_Assert(GtpSend(sock, pkt) == STATUS_OK, status = STATUS_ERROR,
                "GTP Send fail");

    BufblkFree(pkt);

    return status;
}

Status GtpHandleEchoResponse(void *data) {
    UTLT_Assert(data, return STATUS_ERROR, "GTP data is NULL");

    Gtpv1Header *gtpHdr = data;
    UTLT_Assert(gtpHdr->type == GTPV1_ECHO_RESPONSE, return STATUS_ERROR, 
                "The type of GTP data is not 'Echo Response'");

    // TODO: Check the peer device exists, and ....
    // 29.281 says the restart conter shall be ignore by the receiver

    return STATUS_OK;
}

Status GtpHandleEndMark(Sock *sock, void *data) {
    UTLT_Assert(sock && data, return STATUS_ERROR, "GTP data is NULL");
    Status status = STATUS_ERROR;

    // TODO

    status = STATUS_OK;
    return status;
}

Status GtpHandleTPDU(Sock *sock, void *data) {
    UTLT_Assert(sock && data, return STATUS_ERROR, "GTP data is NULL");
    Status status = STATUS_ERROR;
    
    // TODO

    status = STATUS_OK;
    return status;
}
