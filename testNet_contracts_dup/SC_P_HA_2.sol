pragma solidity >0.6.0 <=0.7.0;
pragma experimental ABIEncoderV2;
import "./SC_P_HA_1.sol";

contract SC_P_HA_2
{

		function hospital_leave_generateEstimatedCostBill(address System_Users_Info_address,address P_HA_1_address,uint _pID,uint _estimatedBillID) public
		{
				uint _hID;
				System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
				SC_P_HA_1 P_HA_1_instance=SC_P_HA_1(P_HA_1_address);
				(,_hID,)=System_Users_Info_instance.getHospitalbyAddress(msg.sender);

				SC_P_HA_1.EstimatedCheckUpCost memory estimatedCheckUpCost_Object=P_HA_1_instance.getEstimatedCheckUpCostStruct(_estimatedBillID);

				require(estimatedCheckUpCost_Object.estimatedCostBillID==_estimatedBillID);
				require(estimatedCheckUpCost_Object.hID==_hID);
				require(estimatedCheckUpCost_Object.pID==_pID);
				require(estimatedCheckUpCost_Object.timestampOfLockingByPatient==0);
				require(now-estimatedCheckUpCost_Object.timestampOfLockingByHospital>P_HA_1_instance.timeLimitForPatientLocking());
				require(estimatedCheckUpCost_Object.timestampOfUnlockingByPatient==0);
				require(estimatedCheckUpCost_Object.timestampOfUnlockingByHospital==0);
				estimatedCheckUpCost_Object.timestampOfUnlockingByHospital=now;
				P_HA_1_instance.transferEtherFromThisContract(msg.sender,estimatedCheckUpCost_Object.estimatedCost);
				P_HA_1_instance.setEstimatedCheckUpCostStruct(estimatedCheckUpCost_Object,_estimatedBillID);


		}


		function patient_leave_lockEstimatedAmount(address System_Users_Info_address,address payable P_HA_1_address,uint _hID,uint _estimatedBillID) public
		{
				uint _pID;
				System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
				SC_P_HA_1 P_HA_1_instance=SC_P_HA_1(P_HA_1_address);
				(,_pID , , , , )=System_Users_Info_instance.getPatientbyAddress(msg.sender);

				SC_P_HA_1.EstimatedCheckUpCost memory estimatedCheckUpCost_Object=P_HA_1_instance.getEstimatedCheckUpCostStruct(_estimatedBillID);

				require(estimatedCheckUpCost_Object.estimatedCostBillID==_estimatedBillID);
				require(estimatedCheckUpCost_Object.hID==_hID);
				require(estimatedCheckUpCost_Object.pID==_pID);
				require(estimatedCheckUpCost_Object.timestampOfCheckUpStart==0);
				require(estimatedCheckUpCost_Object.timestampOfUnlockingByPatient==0);
				require(estimatedCheckUpCost_Object.timestampOfUnlockingByHospital==0);
				require(now-estimatedCheckUpCost_Object.timestampOfLockingByPatient>P_HA_1_instance.timeLimitForStartingOperation());
				estimatedCheckUpCost_Object.timestampOfUnlockingByPatient=now;
				P_HA_1_instance.transferEtherFromThisContract(msg.sender,estimatedCheckUpCost_Object.estimatedCost);
				P_HA_1_instance.setEstimatedCheckUpCostStruct(estimatedCheckUpCost_Object,_estimatedBillID);
		}


		function hospital_leave_keepHashToBlockchain(address System_Users_Info_address,address payable P_HA_1_address,uint _pID,uint _estimatedBillID) public
		{
				uint _hID;
				uint _multiSigID;
				System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
				SC_P_HA_1 P_HA_1_instance=SC_P_HA_1(P_HA_1_address);
				(,_hID,)=System_Users_Info_instance.getHospitalbyAddress(msg.sender);

				SC_P_HA_1.EstimatedCheckUpCost memory estimatedCheckUpCost_Object=P_HA_1_instance.getEstimatedCheckUpCostStruct(_estimatedBillID);
				_multiSigID=P_HA_1_instance.estimatedCostBillIDToMultiSigOnMedicalDataID(_estimatedBillID);

				SC_P_HA_1.MultiSigOnMedicalData memory multiSigOnMedicalData_Object=P_HA_1_instance.getMultiSigOnMedicalDataStruct(_multiSigID);


				require(multiSigOnMedicalData_Object.hospitalID==_hID);
				require(multiSigOnMedicalData_Object.patientID==_pID);
				require(multiSigOnMedicalData_Object.timeStampOfSignByHospital!=0 && multiSigOnMedicalData_Object.timeStampOfSignByHospital<now);
				require(multiSigOnMedicalData_Object.timeStampOfUnlockingByHospital_multi==0);
				require(multiSigOnMedicalData_Object.timeStampOfVerificationByPatient==0);
				require(now - multiSigOnMedicalData_Object.timeStampOfSignByHospital>P_HA_1_instance.timeLimitForVerifyingSignature());
				multiSigOnMedicalData_Object.timeStampOfUnlockingByHospital_multi=now;
				P_HA_1_instance.transferEtherFromThisContract(msg.sender,estimatedCheckUpCost_Object.estimatedCost);
				P_HA_1_instance.setMultiSigOnMedicalData(multiSigOnMedicalData_Object,_multiSigID);

		}

		function patient_leave_verifyAndGiveConsent(address System_Users_Info_address,address payable P_HA_1_address,uint _hID,uint _estimatedBillID) public
		{
				uint _pID;
				uint _multiSigID;
				uint _finalBillID;
				System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
				SC_P_HA_1 P_HA_1_instance=SC_P_HA_1(P_HA_1_address);
				(,_pID , , , , )=System_Users_Info_instance.getPatientbyAddress(msg.sender);


				_multiSigID=P_HA_1_instance.estimatedCostBillIDToMultiSigOnMedicalDataID(_estimatedBillID);
				_finalBillID=P_HA_1_instance.estimatedBillIDToFinalBillID(_estimatedBillID);

				SC_P_HA_1.MultiSigOnMedicalData memory multiSigOnMedicalData_Object=P_HA_1_instance.getMultiSigOnMedicalDataStruct(_multiSigID);
				SC_P_HA_1.EstimatedCheckUpCost memory estimatedCheckUpCost_Object=P_HA_1_instance.getEstimatedCheckUpCostStruct(_estimatedBillID);


				require(multiSigOnMedicalData_Object.hospitalID==_hID);
				require(multiSigOnMedicalData_Object.patientID==_pID);
				require(multiSigOnMedicalData_Object.timeStampOfVerificationByPatient!=0 && multiSigOnMedicalData_Object.timeStampOfVerificationByPatient<now);
				require(multiSigOnMedicalData_Object.timeStampOfUnlockingByPatient_multi==0);
				require(_finalBillID==0);




				require(now - multiSigOnMedicalData_Object.timeStampOfVerificationByPatient>P_HA_1_instance.timeLimitForGeneratingFillBill());
				multiSigOnMedicalData_Object.timeStampOfUnlockingByPatient_multi=now;
				P_HA_1_instance.transferEtherFromThisContract(msg.sender,estimatedCheckUpCost_Object.estimatedCost);
				P_HA_1_instance.setMultiSigOnMedicalData(multiSigOnMedicalData_Object,_multiSigID);
		}

		function hospital_leave_DischargeAndGenerateFinalCostBill(address System_Users_Info_address,address payable  P_HA_1_address,uint _estimatedBillID,uint _pID) public
		{
			 uint _hID;
			 uint _finalBillID;
			 System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
			 SC_P_HA_1 P_HA_1_instance=SC_P_HA_1(P_HA_1_address);
			 (,_hID,)=System_Users_Info_instance.getHospitalbyAddress(msg.sender);

			 _finalBillID=P_HA_1_instance.estimatedBillIDToFinalBillID(_estimatedBillID);

			 SC_P_HA_1.FinalCheckUpCost memory finalCheckUpCost_Object=P_HA_1_instance.getFinalBillStruct(_finalBillID);


			 require(finalCheckUpCost_Object.hID==_hID);
			 require(finalCheckUpCost_Object.pID==_pID);
			 require(finalCheckUpCost_Object.timeStampOfUnlockingByHospital_final==0);
			 require(finalCheckUpCost_Object.timestampOfFinalBilling!=0 && finalCheckUpCost_Object.timestampOfFinalBilling<now);
			 require(finalCheckUpCost_Object.timeStampOfPatientFinalBillConsent==0);
			 require(now-finalCheckUpCost_Object.timestampOfFinalBilling>P_HA_1_instance.timeLimitForSigningFinalBill());
			 finalCheckUpCost_Object.timeStampOfUnlockingByHospital_final=now;
			 P_HA_1_instance.transferEtherFromThisContract(msg.sender,finalCheckUpCost_Object.finalCost);
			 P_HA_1_instance.setFinalCheckUpCost(finalCheckUpCost_Object,_finalBillID);

		}

		function patient_leave_consentFinalBill(address System_Users_Info_address,address payable P_HA_1_address,uint _estimatedBillID,uint _hID) public
		{
				uint _pID;
				uint _finalBillID;
				uint _multiSigID;
				System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
				SC_P_HA_1 P_HA_1_instance=SC_P_HA_1(P_HA_1_address);
				(,_pID , , , , )=System_Users_Info_instance.getPatientbyAddress(msg.sender);

				_finalBillID=P_HA_1_instance.estimatedBillIDToFinalBillID(_estimatedBillID);
				_multiSigID=P_HA_1_instance.estimatedCostBillIDToMultiSigOnMedicalDataID(_estimatedBillID);

				SC_P_HA_1.FinalCheckUpCost memory finalCheckUpCost_Object=P_HA_1_instance.getFinalBillStruct(_finalBillID);
				SC_P_HA_1.KeyAndExchange memory keyAndExchange_Object=P_HA_1_instance.getKeyAndExchangeStruct(_finalBillID);


				require(finalCheckUpCost_Object.hID==_hID);
				require(finalCheckUpCost_Object.pID==_pID);
				require(finalCheckUpCost_Object.timeStampOfPatientFinalBillConsent!=0 && finalCheckUpCost_Object.timeStampOfPatientFinalBillConsent<now);
				require(finalCheckUpCost_Object.timestampOfUnlockingByPatient_final==0);
				require(keyAndExchange_Object.timeStampOfKeyReveal==0);
				require(now-finalCheckUpCost_Object.timeStampOfPatientFinalBillConsent>P_HA_1_instance.timeLimitForRaisingDispute());
				finalCheckUpCost_Object.timestampOfUnlockingByPatient_final=now;
				P_HA_1_instance.transferEtherFromThisContract(msg.sender,finalCheckUpCost_Object.finalCost);
 			 	P_HA_1_instance.setFinalCheckUpCost(finalCheckUpCost_Object,_finalBillID);

		}



		function hospital_leave_keyReveal(address System_Users_Info_address,address payable P_HA_1_address,uint _estimatedBillID,uint _pId) public
		{
				uint _hID;
				uint _finalBillID;
				uint _multiSigID;
				System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
				SC_P_HA_1 P_HA_1_instance=SC_P_HA_1(P_HA_1_address);
				(,_hID,)=System_Users_Info_instance.getHospitalbyAddress(msg.sender);

				_finalBillID=P_HA_1_instance.estimatedBillIDToFinalBillID(_estimatedBillID);
				_multiSigID=P_HA_1_instance.estimatedCostBillIDToMultiSigOnMedicalDataID(_estimatedBillID);

				SC_P_HA_1.FinalCheckUpCost memory finalCheckUpCost_Object=P_HA_1_instance.getFinalBillStruct(_finalBillID);
				SC_P_HA_1.KeyAndExchange memory keyAndExchange_Object=P_HA_1_instance.getKeyAndExchangeStruct(_finalBillID);


				require(finalCheckUpCost_Object.hID==_hID);
				require(finalCheckUpCost_Object.pID==	_pId);
				require(finalCheckUpCost_Object.timeStampOfUnlockingByHospital_final==0);
				require(keyAndExchange_Object.timeStampOfKeyReveal!=0 && keyAndExchange_Object.timeStampOfKeyReveal<now);
				require(keyAndExchange_Object.timeStampOfComplainOrAcceptance==0);
				require(now-keyAndExchange_Object.timeStampOfKeyReveal>P_HA_1_instance.timeLimitForAcceptingOrProducingComplain());
				finalCheckUpCost_Object.timeStampOfFinalConsent=now;
				P_HA_1_instance.transferEtherFromThisContract(msg.sender,finalCheckUpCost_Object.finalCost);
				P_HA_1_instance.setFinalCheckUpCost(finalCheckUpCost_Object,_finalBillID);

		}


}
