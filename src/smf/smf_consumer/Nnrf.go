package smf_consumer

import (
	"context"
	"github.com/antihax/optional"
	"github.com/mohae/deepcopy"
	"free5gc/lib/Nnrf_NFDiscovery"
	"free5gc/lib/openapi/models"
	"free5gc/src/smf/logger"
	"free5gc/src/smf/smf_context"
	"net/http"
)

func SendNFRegistration() {

	//set nfProfile
	profile := models.NfProfile{
		NfInstanceId:  smf_context.SMF_Self().NfInstanceID,
		NfType:        models.NfType_SMF,
		NfStatus:      models.NfStatus_REGISTERED,
		Ipv4Addresses: []string{smf_context.SMF_Self().HTTPAddress},
		NfServices:    smf_context.NFServices,
		SmfInfo:       smf_context.SmfInfo,
	}

	// Check data (Use RESTful PUT)
	rep, res, err := smf_context.SMF_Self().NFManagementClient.NFInstanceIDDocumentApi.RegisterNFInstance(context.TODO(), smf_context.SMF_Self().NfInstanceID, profile)

	if err != nil {
		logger.AppLog.Panic(err)
	}

	if res != nil {
		if status := res.StatusCode; status != http.StatusOK {
			if status != http.StatusCreated {
				logger.AppLog.Info("handler returned wrong status code", status)
			}
		}
	}

	logger.InitLog.Infof("SMF Registration to NRF %v", rep)

}

func SendNFDeregistration() {

	// Check data (Use RESTful DELETE)
	res, err := smf_context.SMF_Self().NFManagementClient.NFInstanceIDDocumentApi.DeregisterNFInstance(context.TODO(), smf_context.SMF_Self().NfInstanceID)
	if err != nil {
		logger.AppLog.Panic(err)
	}
	if res != nil {
		if status := res.StatusCode; status != http.StatusNoContent {
			logger.AppLog.Info("handler returned wrong status code", status)
		}
	}

}

func SendNFDiscoveryUDM() {
	// Set targetNfType
	targetNfType := models.NfType_UDM
	// Set requestNfType
	requesterNfType := models.NfType_SMF
	localVarOptionals := Nnrf_NFDiscovery.SearchNFInstancesParamOpts{}

	// Check data
	rep, res, err := smf_context.SMF_Self().NFDiscoveryClient.NFInstancesStoreApi.SearchNFInstances(context.TODO(), targetNfType, requesterNfType, &localVarOptionals)
	if err != nil {
		logger.AppLog.Panic(err)
	}
	if res != nil {
		if status := res.StatusCode; status != http.StatusOK {
			logger.AppLog.Info("handler returned wrong status code", status)
		}
	}

	smf_context.SMF_Self().UDMProfiles = rep.NfInstances
	//logger.AppLog.Info(smf_context.UDMNFProfile)
}

func SendNFDiscoveryPCF() {

	// Set targetNfType
	targetNfType := models.NfType_PCF
	// Set requestNfType
	requesterNfType := models.NfType_SMF
	localVarOptionals := Nnrf_NFDiscovery.SearchNFInstancesParamOpts{}

	// Check data
	rep, res, err := smf_context.SMF_Self().NFDiscoveryClient.NFInstancesStoreApi.SearchNFInstances(context.TODO(), targetNfType, requesterNfType, &localVarOptionals)
	if err != nil {
		logger.AppLog.Panic(err)
	}
	if res != nil {
		if status := res.StatusCode; status != http.StatusOK {
			logger.AppLog.Info("handler returned wrong status code", status)
		}
	}

	smf_context.SMF_Self().PCFProfiles = rep.NfInstances
	//logger.AppLog.Info(smf_context.PCFNFProfile)
}

func SendNFDiscoveryServingAMF(smContext *smf_context.SMContext) {
	targetNfType := models.NfType_AMF
	requesterNfType := models.NfType_SMF

	localVarOptionals := Nnrf_NFDiscovery.SearchNFInstancesParamOpts{}

	localVarOptionals.TargetNfInstanceId = optional.NewInterface(smContext.ServingNfId)

	// Check data
	rep, res, err := smf_context.SMF_Self().NFDiscoveryClient.NFInstancesStoreApi.SearchNFInstances(context.TODO(), targetNfType, requesterNfType, &localVarOptionals)
	if err != nil {
		logger.AppLog.Panic(err)
		return
	}
	if rep.NfInstances == nil {
		if status := res.StatusCode; status != http.StatusOK {
			logger.AppLog.Info("handler returned wrong status code", status)
		}
		logger.AppLog.Info("rep.NfInstances == nil")
		return
	}
	logger.AppLog.Info("SendNFDiscoveryServingAMF ok")
	smContext.AMFProfile = deepcopy.Copy(rep.NfInstances[0]).(models.NfProfile)

}
