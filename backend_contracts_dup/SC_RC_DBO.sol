pragma solidity >0.4.0 <=0.7.0;
pragma experimental ABIEncoderV2;
/* import "./p_ha_11.sol"; */


contract SC_RC_DBO
{

    uint public Now;
    address govtAuthority;

    struct ResearchDataDetails
    {
        uint id;
        address rcAddr;
        address dataBaseOwner;
        uint timeStamp_researchCommunityRequestData;
        uint timeStamp_dataBaseOwnerSendData;
        bytes32 hashOfMedData;
        bytes sigOnHashOfMedData;

    }


    function increaseTheTimeBy10seconds() public {
			Now=now;
		}

    ResearchDataDetails[] public researchDataDetails;


    uint public researchDataDetailsId=0;
    uint public researchCommunityId=0;
    uint public timeLimitForGivingData=50;



    mapping(uint=>address) public researchCommunityIdToAddress;
    mapping(address=>mapping(address=>uint)) public researchCommunity_databaseOwner_researchDataId;



    function requestMedicalData(address System_Users_Info_address,address _dataBaseOwner) public
    {

        System_Users_Info System_Users_Info_instance = System_Users_Info(System_Users_Info_address);
        (_rcAddress,,)=System_Users_Info_instance.rcById(_researchCommunityId);

        require(System_Users_Info_instance.rcAddressTorcId(msg.sender)!=0);
        require(System_Users_Info_instance,dboAddressTodboId(msg.sender)!=0);

        researchDataDetailsId+=1;
        ResearchDataDetails memory ResearchDataDetailsInstance;
        ResearchDataDetailsInstance.id=researchDataDetailsId;
        ResearchDataDetailsInstance.rcAddr=msg.sender;
        ResearchDataDetailsInstance.dataBaseOwner=_dataBaseOwner;
	      ResearchDataDetailsInstance.timeStamp_researchCommunityRequestData=now;
        researchDataDetails.push(ResearchDataDetailsInstance);
        researchCommunity_databaseOwner_researchDataId[msg.sender][_dataBaseOwner]=researchDataDetailsId;

    }


    function provideMedicalData(address System_Users_Info_address,uint _researchDataId,bytes32 _hashOfMedData,bytes memory _sigOnHashOfMedData) public
    {
        require(researchDataDetails[_researchDataId-1].dataBaseOwner==msg.sender);
        require(now-researchDataDetails[_researchDataId-1].timeStamp_researchCommunityRequestData<=timeLimitForGivingData);

        researchDataDetails[_researchDataId-1].hashOfMedData=_hashOfMedData;
        researchDataDetails[_researchDataId-1].sigOnHashOfMedData=_sigOnHashOfMedData;
        researchDataDetails[_researchDataId-1].timeStamp_dataBaseOwnerSendData=now;

    }


}
