pragma solidity >0.4.0 <=0.7.0;
pragma experimental ABIEncoderV2;
import "./SC_P_HA_1.sol";
import "./SC_P_DBO.sol";


contract SC_P_IC_DBO
{


    struct policyBuyPhaseOne
    {
        uint p_price;  //price locked by buyerAddr..
        uint time;    // time of money locking..
        bytes32 hash_of_file; //hash of the t&c file..
    }

    struct PolicyDetails
    {
        uint pid; //policy ID..
        address payable buyerAddr; //buyerAddr address..
        address payable icAddr; //insurance company address..
        uint timestamp; //timestamp..
        uint policyPrice; //price of the policy..
        bytes32 hashOfTermsAndCon; //hash of terms&con file..
        uint[] claimIDs;
    }

    struct ClaimDetails
    {
        uint cID; //claim
        uint estimatedBillID;
        uint timestamp_patientGenerateClaim; //timestamp_policyBuying..
        uint claimedAmount; //claimed amount..\
        uint approvedAmount;
        bytes32 comm_k;
        bytes32 K;
        uint timestamp_keyReveal;
        uint timestamp_icLocking;
        uint timestamp_icUnlocking;
        uint timestamp_approval;
        bool isSelfApproved;
        bool isICMalicious;
    }

    struct DBODetails{
      address dboAddress;
      bytes signature_DBO_hashOfEncryptedFile;
      uint timeStamp_DBgivingEncryptedFile;
    }

    struct ICReputation
    {
    	uint noOfCustomers;
    	uint noOfMaliciousTagReceived;
    }


    PolicyDetails[] public policyDetails;
    uint public pidGenerator=0;

    ClaimDetails[] public claimDetails;
    uint public cidGenerator=0;



    function returnClaimIds(uint _pId) public view returns(uint[] memory)
    {
      return policyDetails[_pId-1].claimIDs;
    }



    mapping(address=>ICReputation) public reputationOfIC;
    mapping(uint=>DBODetails) public claimIDToDBODetails;
    mapping(uint=>uint) cidTopid;
    mapping(address=>mapping(address=>policyBuyPhaseOne)) public policyPhaseOne;
    mapping(address=>mapping(address=>uint)) public policyId_map;
    mapping(address=>mapping(address=>bool)) public oneTimeKey_isExpired;


    mapping(address=>uint) public securityMoney_IC;
    mapping(address=>bool) public deRegister_IC;

    uint constant timeLimitForBuying = 100 seconds ;
    uint constant timeLimitForFinalApproval = 100 seconds;
    uint constant timeLimitForIcLocking=100 seconds;
    uint constant timeLimitForKeyReveal=100 seconds;



    function addSecurityMoney(address System_Users_Info_address,uint _amount) public payable
    {
        System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
      	require(System_Users_Info_instance.icAddressToicId(msg.sender)!=0); //checking for valid ic id
        require(msg.value>=_amount);
      	securityMoney_IC[msg.sender] += _amount;
    }

    function withdrawSecurityMoney(address System_Users_Info_address,uint _amount) public
    {
        System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
        require(System_Users_Info_instance.icAddressToicId(msg.sender)!=0); //checking for valid ic id
        require(securityMoney_IC[msg.sender]-_amount>=0);
      	securityMoney_IC[msg.sender] -= _amount;
      	msg.sender.transfer(_amount);
    }



    function buyPolicyPhaseOne(address System_Users_Info_address,uint _icID, bytes32 _hashOfTermsAndCon) public payable
    {
        address payable _icAddr;
        System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
        (_icAddr, , )=System_Users_Info_instance.getInsuranceCompanybyID(_icID);
        policyBuyPhaseOne memory x;
        x=policyPhaseOne[msg.sender][_icAddr];
        require(x.p_price==0);
        x.p_price=msg.value;
        x.time=now;
        x.hash_of_file=_hashOfTermsAndCon;
        policyPhaseOne[msg.sender][_icAddr] = x;

    }

    /**
     *  Now, the policy buyerAddr will communicate offline to the Insurance Company for providing the hash
     * of the terms&condition file(intended policy) within some fixed time window(i.e timeout). Else the buyerAddr can reclaim
     * the locked money..
     **/

    //The following function should be invoked by the insurance company..
    function buyPolicyPhaseTwo(address System_Users_Info_address,uint _patientId, uint _price, bytes32 _hashOfTermsAndCon) public
    {
        address payable _buyerAddr;
        System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
        (_buyerAddr, , , , , )=System_Users_Info_instance.getPatientbyID(_patientId);


        policyBuyPhaseOne memory x;
        x=policyPhaseOne[_buyerAddr][msg.sender];
        require(_hashOfTermsAndCon == x.hash_of_file);
        require(_price == x.p_price);
        require(now - x.time <= timeLimitForBuying);

        uint _securityMoney=securityMoney_IC[msg.sender];
        if(_price>_securityMoney)
        {
          deRegister_IC[msg.sender]=true;
          revert();
        }


        pidGenerator += 1;
        PolicyDetails memory pd;
        pd.pid=pidGenerator;
        pd.buyerAddr=_buyerAddr;
        pd.icAddr=msg.sender;
        pd.timestamp=now;
        pd.policyPrice=_price;
        pd.hashOfTermsAndCon=_hashOfTermsAndCon;
        policyDetails.push(pd);
        policyPhaseOne[_buyerAddr][msg.sender].p_price=0;
        policyId_map[_buyerAddr][msg.sender]=pidGenerator;
        msg.sender.transfer(_price);

        reputationOfIC[msg.sender].noOfCustomers+=1;

    }


    function patient_withdrawLockedPolicyBuyMoney(address System_Users_Info_address,uint _icID) public
    {
        address _ic;
        System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
        (_ic, , )=System_Users_Info_instance.getInsuranceCompanybyID(_icID);
        policyBuyPhaseOne memory temp = policyPhaseOne[msg.sender][_ic];

        require(now > temp.time+timeLimitForBuying);
        require(policyPhaseOne[msg.sender][_ic].p_price!=0);
        msg.sender.transfer(policyPhaseOne[msg.sender][_ic].p_price);
        temp.p_price=0;

    }


    function claimMoney(address System_Users_Info_address,address SC_P_HA_1_address,address SC_P_DBO_address,uint _policyID, uint _claimedAmount,uint _estimatedBillId,bytes memory _publicKey, bytes32 _comm_k,address _dboAddr,uint _appID) public
    {


        uint icId;

        System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);



        claimMoneyCheckConditions(System_Users_Info_address,SC_P_HA_1_address,SC_P_DBO_address,_estimatedBillId,_policyID,_appID,_dboAddr,msg.sender,_claimedAmount);



        ClaimDetails memory cd;
        PolicyDetails storage pd=policyDetails[_policyID-1];

        for(uint i=0;i<pd.claimIDs.length;i++)
        {
          cd=claimDetails[pd.claimIDs[i]-1];

          if(cd.estimatedBillId==_estimatedBillId)
          {
            require(cd.timestamp_Approval==0);
            require(now-cd.timestamp_patientGenerateClaim>timeLimitForIcLocking);
          }

        }

        cidGenerator = cidGenerator+1;


        (,icId,)=System_Users_Info_instance.getInsuranceCompanybyAddress(pd.icAddr);


        cd=ClaimDetails(cidGenerator,_estimatedBillId,now,_claimedAmount,0,_comm_k,'',0,0,0,0,false,false);
        claimDetails.push(cd);

        DBODetails memory dboDetails;

        dboDetails.dboAddress=_dboAddr;
        claimIDToDBODetails[cidGenerator]=dboDetails;

        System_Users_Info_instance.grant_permission_patient_To_IC(msg.sender,icId);
        pd.claimIDs.push(cidGenerator);
        cidTopid[cidGenerator]=_policyID;


      }

      function claimMoneyCheckConditions(address System_Users_Info_address,address SC_P_HA_1_address,address SC_P_DBO_address,uint _estimatedBillId,uint _policyID,uint _appID,address _dboAddr,address _fnCallerAdress,uint _claimedAmount) private
      {

          uint _multiSigID;
          uint _finalBillID;


          System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
          SC_P_DBO SC_P_DBO_instance = SC_P_DBO(SC_P_DBO_address);
          SC_P_HA_1 SC_P_HA_1_instance=SC_P_HA_1(SC_P_HA_1_address);


          _multiSigID=SC_P_HA_1_instance.estimatedCostBillIDToMultiSigOnMedicalDataID(_estimatedBillId);
          _finalBillID=SC_P_HA_1_instance.estimatedBillIDToFinalBillID(_estimatedBillId);

          SC_P_DBO.ApplicationForStoring memory ApplicationForStoringStruct =SC_P_DBO_instance.getApplicationForStoringStruct(_appID);

          SC_P_HA_1.FinalCheckUpCost memory FinalCheckUpCostStruct = SC_P_HA_1_instance.getFinalBillStruct(_finalBillID);

          require(FinalCheckUpCostStruct.finalCost>=_claimedAmount);
          require(policyDetails[_policyID-1].buyerAddr==_fnCallerAdress);
          require(ApplicationForStoringStruct.patientAddress==_fnCallerAdress);
          require(ApplicationForStoringStruct.dboAddress==_dboAddr);
          require(ApplicationForStoringStruct.multiSigID==_multiSigID);
          require(ApplicationForStoringStruct.timestamp_approval!=0);


      }

      function dboKeepSigOnHashOfEncFile(uint _claimId,bytes memory _signature_DBO_hashOfEncryptedFile) public
      {

          require(claimIDToDBODetails[_claimId].dboAddress==msg.sender);
          require(claimDetails[_claimId-1].timestamp_IcLocking==0);
          require(claimDetails[_claimId-1].timestamp_patientGenerateClaim!=0);

          claimIDToDBODetails[_claimId].signature_DBO_hashOfEncryptedFile=_signature_DBO_hashOfEncryptedFile;
          claimIDToDBODetails[_claimId].timeStamp_DBgivingEncryptedFile=now;


          uint _pId;
          address payable _buyerAddr;
          address payable _icAddr;
          _pId=cidTopid[_claimId];
          _buyerAddr=policyDetails[_pId-1].buyerAddr;
          _icAddr=policyDetails[_pId-1].icAddr;
          oneTimeKey_isExpired[_buyerAddr][_icAddr]=true;

      }


      function lockClaimedMoney(address System_Users_Info_address,uint _cID,uint _patientId) public payable
      {

          address _buyerAddr;

          System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
          (_buyerAddr, , , , , )=System_Users_Info_instance.getPatientbyID(_patientId);

           uint _pid;
          _pid =cidTopid[_cID];
          require(policyDetails[_pid-1].icAddr == msg.sender);
          require(policyDetails[_pid-1].buyerAddr==_buyerAddr);


          require(now-claimDetails[_cID-1].timestamp_patientGenerateClaim<=timeLimitForIcLocking);
          require(claimDetails[_cID-1].timestamp_IcLocking==0);
          require(msg.value>=claimDetails[_cID-1].claimedAmount);

          claimDetails[_cID-1].timestamp_IcLocking=now;

      }

      function returnPolicyPrice(address System_Users_Info_address,uint _cID,uint _icID) public
      {


      	  address _icAddr;
      	  System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
      	  (_icAddr, , )=System_Users_Info_instance.getInsuranceCompanybyID(_icID);

      	  uint _securityMoney=securityMoney_IC[_icAddr];
      	  uint _poID;
          _poID =cidTopid[_cID];

          require(policyDetails[_poID-1].icAddr == _icAddr);
          require(policyDetails[_poID-1].buyerAddr == msg.sender);
          require(claimDetails[_cID-1].timestamp_IcLocking==0);
          require(now-claimDetails[_cID-1].timestamp_patientGenerateClaim > timeLimitForIcLocking);


          claimDetails[_cID-1].isICMalicious=true;
          reputationOfIC[msg.sender].noOfMaliciousTagReceived += 1;


          if(policyDetails[_poID-1].policyPrice > _securityMoney) // cancel IC Registration..
          {
              	deRegister_IC[msg.sender]=true;
              	msg.sender.transfer(_securityMoney); //partial payment done..
              	_securityMoney=0;
          }
          else
          {
          	msg.sender.transfer(policyDetails[_poID-1].policyPrice); //full policy price is returned..
          	_securityMoney -= policyDetails[_poID-1].policyPrice;
          }

          securityMoney_IC[_icAddr]=_securityMoney;

      }





      function withdrawingLockedClaimedMoney(address System_Users_Info_address,uint _cID,uint _patientId) public
      {

          address _buyerAddr;

          System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
          (_buyerAddr, , , , , )=System_Users_Info_instance.getPatientbyID(_patientId);

          uint _pid;
          _pid =cidTopid[_cID]; //getting the polcy no for which claim is made..
          require(policyDetails[_pid-1].icAddr == msg.sender);
          require(policyDetails[_pid-1].buyerAddr==_buyerAddr);


          require(claimDetails[_cID-1].timestamp_IcLocking!=0);
          require(now-claimDetails[_cID-1].timestamp_IcLocking>timeLimitForKeyReveal);
          require(claimDetails[_cID-1].timestamp_keyReveal==0);
          claimDetails[_cID-1].timestamp_IcUnlocking=now;
          msg.sender.transfer(claimDetails[_cID-1].claimedAmount);

      }



      function revealSecretKey(address System_Users_Info_address,bytes32 _K,uint _cID,uint _icID) public
      {
          address _icAddr;
          System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
          (_icAddr, , )=System_Users_Info_instance.getInsuranceCompanybyID(_icID);

          uint _pid;
          _pid =cidTopid[_cID]; //getting the polcy no for which claim is made..
          require(policyDetails[_pid-1].icAddr == _icAddr);
          require(policyDetails[_pid-1].buyerAddr==msg.sender);



          require(now-claimDetails[_cID-1].timestamp_IcLocking<=timeLimitForKeyReveal);
          require(claimDetails[_cID-1].timestamp_keyReveal==0);
          require(claimDetails[_cID-1].timestamp_Approval==0);
          require(Hash(_K)==claimDetails[_cID-1].comm_k);

          claimDetails[_cID-1].secretKey=_secretKey;
          claimDetails[_cID-1].timestamp_keyReveal=now;

      }




      function approveClaim(address System_Users_Info_address,uint _cID,uint _patientId,uint _approvedAmount) payable public
      {

            address payable _buyerAddr;

            System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
            (_buyerAddr, , , , , )=System_Users_Info_instance.getPatientbyID(_patientId);

            uint _pid;
            _pid =cidTopid[_cID]; //getting the polcy no for which claim is made..
            require(policyDetails[_pid-1].icAddr == msg.sender);
            require(policyDetails[_pid-1].buyerAddr==_buyerAddr);


            require(now-claimDetails[_cID-1].timestamp_keyReveal<=timeLimitForFinalApproval);
            require(claimDetails[_cID-1].timestamp_Approval==0);
            require(claimDetails[_cID-1].timestamp_keyReveal!=0);
            require(msg.value>=_approvedAmount);

            claimDetails[_cID-1].approvedAmount=_approvedAmount;
            claimDetails[_cID-1].timestamp_Approval=now;
            /* _buyerAddr.transfer(claimDetails[_cID-1].approvedAmount); */


      }


      function selfApproveClaim(address System_Users_Info_address,uint _cID,uint _icID) public
      {
          address _ic;
          System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
          (_ic, , )=System_Users_Info_instance.getInsuranceCompanybyID(_icID);

          uint _pid;
          _pid =cidTopid[_cID]; //getting the polcy no for which claim is made..
          require(policyDetails[_pid-1].icAddr == _ic);
          require(policyDetails[_pid-1].buyerAddr==msg.sender);

          require(now-claimDetails[_cID-1].timestamp_keyReveal>timeLimitForFinalApproval);
          require(claimDetails[_cID-1].timestamp_Approval==0);
          require(claimDetails[_cID-1].timestamp_keyReveal!=0);


          claimDetails[_cID-1].timestamp_Approval=now;
          claimDetails[_cID-1].isSelfApproved=true;
          msg.sender.transfer(claimDetails[_cID-1].claimedAmount);

      }


      function Hash(bytes memory _toBeHashed) private pure returns(bytes32)
	    {
    			bytes32 Hashed;
    			bytes memory toBeHashed;
    			toBeHashed=new bytes (_toBeHashed.length);
    			for (uint i=0;i<_toBeHashed.length;i++)
    			{
    				toBeHashed[i]=_toBeHashed[i];
    			}
    			Hashed=keccak256(toBeHashed);
    			return Hashed;

	     }


}
