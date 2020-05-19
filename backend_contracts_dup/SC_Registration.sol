pragma solidity >0.6.0 <=0.7.0;
pragma experimental ABIEncoderV2;

contract System_Users_Info
{

    address public govtAuthority;


    constructor() public {
      govtAuthority=msg.sender;
    }


    modifier onlyGovtAuthority(address _party)
    {
        require(_party==govtAuthority);
        _;
    }



    /* Patient Registration */

    uint public patientIDGenerator=0;

    struct Patient
    {
        address payable paddr;
        uint patientID;
        string name;
        uint8 age;
        uint mob;
        string add;
    }

    Patient[] public patient;

    mapping(address=>uint) public mapFromAddToId_P; //Mapping from patient address to patient ID..

    function patientRegistration(string memory _name, uint _age, uint _mob, string memory _add) public
    {
        require(mapFromAddToId_P[msg.sender] == 0); //The user was already registered..
        require(!isEmptyString(_name)); //Checking whether name is empty or not..
        require(_age>0 && _age<=100); //Checking for valid age range..
        require(_mob>0 && numDigits(_mob)==10); //Checking for valid mobile no..
        require(!isEmptyString(_add)); //Checking whether address is empty or not..

        //The user will be registered..
        patientIDGenerator = patientIDGenerator + 1;
        mapFromAddToId_P[msg.sender]=patientIDGenerator;
        Patient memory pt = Patient(msg.sender,patientIDGenerator,_name,uint8(_age),_mob,_add);
        patient.push(pt);

    }


    function getPatientbyID(uint _patientID) view public returns(address payable,uint,string memory,uint8,uint,string memory)
    {
        require(_patientID>=1 && _patientID<=patientIDGenerator);
        uint x = _patientID-1;
        return(patient[x].paddr,patient[x].patientID,patient[x].name,patient[x].age,patient[x].mob,patient[x].add);
    }

    function getPatientbyAddress(address _patientAdd) view public returns(address payable,uint,string memory,uint8,uint,string memory)
    {
        uint _patientID = mapFromAddToId_P[_patientAdd];
        require(_patientID>=1 && _patientID<=patientIDGenerator);
        uint x = _patientID-1;
        return(patient[x].paddr,patient[x].patientID,patient[x].name,patient[x].age,patient[x].mob,patient[x].add);

    }






    /*Hospital Registration*/

    uint public hospitalIDGenerator=0;

    struct Hospital
    {
        address payable haddr;
        uint hospitalID;
        string name;
    }

    Hospital[] public hospital;

    mapping(address=>uint) mapFromAddToId_H; //Mapping from hospital address to hospital ID..

    function hospitalRegistration(string memory _name) public
    {

        require(mapFromAddToId_H[msg.sender] == 0); //this Hospital was already registered..
        require(!isEmptyString(_name)); //Checking whether name is empty or not..

        //Hospital will be registered..
        hospitalIDGenerator = hospitalIDGenerator + 1;
        mapFromAddToId_H[msg.sender]=hospitalIDGenerator;
        Hospital memory h = Hospital(msg.sender,hospitalIDGenerator,_name);
        hospital.push(h);
    }

    function getHospitalbyID(uint _hospitalID) view public returns(address payable,uint,string memory)
    {
        require(_hospitalID>=1 && _hospitalID<=hospitalIDGenerator);
        uint x = _hospitalID-1;
        return(hospital[x].haddr,hospital[x].hospitalID,hospital[x].name);
    }

    function getHospitalbyAddress(address _hospitalAdd) view public returns(address payable,uint,string memory)
    {
        uint _hospitalID = mapFromAddToId_H[_hospitalAdd];
        require(_hospitalID>=1 && _hospitalID<=hospitalIDGenerator);
        uint x = _hospitalID-1;
        return(hospital[x].haddr,hospital[x].hospitalID,hospital[x].name);
    }




    /*InsuranceCompany Registration*/

    uint public insuranceIDGenerator=0;

    struct Insurance
    {
        address payable icaddr;
        uint insuranceCompanyID;
        string name;
    }

    Insurance[] public insurance;

    mapping(address=>uint) public icAddressToicId; //Mapping from Insurance address to Insurance ID..


    function insuranceCompanyRegistration(address payable _icaddr,string memory _name) public
    {

        require(icAddressToicId[_icaddr]==0);
        require(!isEmptyString(_name));
        insuranceIDGenerator+=1;

        Insurance memory InsuranceObject;
        InsuranceObject.icaddr=_icaddr;
        InsuranceObject.insuranceCompanyID=insuranceIDGenerator;
        InsuranceObject.name=_name;
        insurance.push(InsuranceObject);

        icAddressToicId[_icaddr]=insuranceIDGenerator;


    }


    function getInsuranceCompanybyID(uint _insuranceID) view public returns(address payable,uint,string memory)
    {
        require(_insuranceID>=1 && _insuranceID<=insuranceIDGenerator);
        uint x = _insuranceID-1;
        return(insurance[x].icaddr,insurance[x].insuranceCompanyID,insurance[x].name);
    }

    function getInsuranceCompanybyAddress(address _insuranceAdd) view public returns(address payable,uint,string memory)
    {
        uint _insuranceID = icAddressToicId[_insuranceAdd];
        require(_insuranceID>=1 && _insuranceID<=insuranceIDGenerator);
        uint x = _insuranceID-1;
        return(insurance[x].icaddr,insurance[x].insuranceCompanyID,insurance[x].name);
    }







    /*DataBase Registration*/

    uint public dboIDGenerator=0;

    struct DataBaseOwner
    {
        address payable dboAddress;
        uint dboID;
        string dboName;
    }


    DataBaseOwner[] public dataBaseOwner;
    mapping(address=>uint) public dboAddressTodboID;


    function dboRegistration(address payable _dboAddress,string memory _dboName) public onlyGovtAuthority(msg.sender)
    {
        require(dboAddressTodboID[_dboAddress]==0);
        require(!isEmptyString(_dboName));
        dboIDGenerator+=1;

        DataBaseOwner memory DataBaseOwnerStruct;
        DataBaseOwnerStruct.dboAddress=_dboAddress;
        DataBaseOwnerStruct.dboID=dboIDGenerator;
        DataBaseOwnerStruct.dboName=_dboName;
        dataBaseOwner.push(DataBaseOwnerStruct);

        dboAddressTodboID[_dboAddress]=dboIDGenerator;

    }

    function dboById(uint _dboID) public view returns(address payable,uint,string memory)
    {
        require(_dboID>=1 && _dboID<=dboIDGenerator);
        return (dataBaseOwner[_dboID-1].dboAddress,dataBaseOwner[_dboID-1].dboID,dataBaseOwner[_dboID-1].dboName);
    }

    function dboByAddress(address _dboAddress) public view returns(address payable,uint,string memory)
    {
        uint _dboID=dboAddressTodboID[_dboAddress];
        require(_dboID>=1 && _dboID<=dboIDGenerator);
        return (dataBaseOwner[_dboID-1].dboAddress,dataBaseOwner[_dboID-1].dboID,dataBaseOwner[_dboID-1].dboName);

    }




    /*ResearchCommunity Registration*/

    uint public rcIDGenerator=0;

    struct ResearchCommunity
    {
        address payable rcAddress;
        uint rcID;
        string rcName;
    }


    ResearchCommunity[] public researchCommunity;
    mapping(address=>uint) public rcAddressTorcID;


    function rcRegistration(address payable _rcAddress,string memory _rcName) onlyGovtAuthority(msg.sender) public
    {
        require(rcAddressTorcID[_rcAddress]==0);
        require(!isEmptyString(_rcName));
        rcIDGenerator+=1;

        ResearchCommunity memory ResearchCommunityStruct;
        ResearchCommunityStruct.rcAddress=_rcAddress;
        ResearchCommunityStruct.rcID=rcIDGenerator;
        ResearchCommunityStruct.rcName=_rcName;

        rcAddressTorcID[_rcAddress]=rcIDGenerator;
    }


    function rcById(uint _rcID) public view returns(address payable,uint,string memory)
    {
        require(_rcID>=1 && _rcID<=rcIDGenerator);
        return(researchCommunity[_rcID-1].rcAddress,researchCommunity[_rcID-1].rcID,researchCommunity[_rcID-1].rcName);
    }


    function rcByAddress(address payable _rcAddress) public view returns(address payable,uint,string memory)
    {
        uint _rcID=rcAddressTorcID[_rcAddress];
        require(_rcID>=1 && _rcID<=rcIDGenerator);
        return(researchCommunity[_rcID-1].rcAddress,researchCommunity[_rcID-1].rcID,researchCommunity[_rcID-1].rcName);

    }








    //Permission Data Structure...
    mapping (address=>mapping(address=>bool)) public read_permission_to_H;
    mapping (address=>mapping(address=>bool)) public write_permission_to_H;
    mapping (address=>mapping(address=>bool)) public read_permission_to_IC;


    function grant_permission(address callingPatientAddress,uint _hID) public
    {
        address payable _hAddr;
        (_hAddr, , )=getHospitalbyID(_hID);
        read_permission_to_H[callingPatientAddress][_hAddr]=true;
        write_permission_to_H[callingPatientAddress][_hAddr]=true;
    }

    function grant_permission_patient_To_IC(address callingPatientAddress,uint _IC) public
    {
        address payable icaddr;
        (icaddr, , )=getInsuranceCompanybyID(_IC);
        read_permission_to_IC[callingPatientAddress][icaddr]=true;
    }


    function deny_permission_patient_To_IC(address callingPatientAddress,uint _IC) public
    {
        address payable icaddr;
        (icaddr, , )=getInsuranceCompanybyID(_IC);
        read_permission_to_IC[callingPatientAddress][icaddr]=false;
    }








		function isEmptyString(string memory s) public pure returns(bool)
    {
            bytes memory b=bytes(s);
            if(b.length==0)
            {
                return true;
            }
            return false;
    }

    function numDigits(uint number) public pure returns (uint8)
    {
        uint8 digits = 0;
        while (number != 0)
        {
            number /= 10;
            digits++;
        }
        return digits;
    }

    function isEqual(string memory a, string memory b) public pure returns(bool)
    {
        bytes memory aa=bytes(a);
        bytes memory bb=bytes(b);
        return(keccak256(aa)==keccak256(bb));
    }


}
