pragma solidity >0.6.0 <=0.7.0;
pragma experimental ABIEncoderV2;

contract System_Users_Info
{
    address public constant databaseOwnerAddr = 0x4B0897b0513fdC7C541B6d9D7E929C4e5364D2dB;

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
        if(mapFromAddToId_P[msg.sender] != 0) //The user was already registered..
        {
            revert();
        }
        //Data Validation..
        if (isEmptyString(_name)) //Checking whether name is empty or not..
        {
            revert();
        }
        if(_age<=0 || _age>=100) //Checking for valid age range..
        {
            revert();
        }
        if(_mob<0 || numDigits(_mob)!=10) //Checking for valid mobile no..
        {
            revert();
        }
        if (isEmptyString(_add)) //Checking whether address is empty or not..
        {
            revert();
        }
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
        require(msg.sender != databaseOwnerAddr);
        if(mapFromAddToId_H[msg.sender] != 0) //this Hospital was already registered..
        {
            revert();
        }
        //Data Validation..
        if (isEmptyString(_name)) //Checking whether name is empty or not..
        {
            revert();
        }
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


    uint public insuranceIDGenerator=0;

    struct Insurance
    {
        address payable icaddr;
        uint insuranceCompanyID;
        string name;
    }

    Insurance[] public insurance;

    mapping(address=>uint) public mapFromAddToId_I; //Mapping from Insurance address to Insurance ID..

    mapping(address=>uint) public securityMoney_IC;
    mapping(address=>bool) public deRegister_IC;

    function insuranceCompanyRegistration(string memory _name) public
    {
        require(msg.sender != databaseOwnerAddr);
        if(mapFromAddToId_I[msg.sender] != 0) //this Insurance Company was already registered..
        {
            revert();
        }
        //Data Validation..
        if (isEmptyString(_name)) //Checking whether name is empty or not..
        {
            revert();
        }
        //Insurance Company will be registered..
        insuranceIDGenerator = insuranceIDGenerator + 1;
        mapFromAddToId_I[msg.sender]=insuranceIDGenerator;
        Insurance memory h = Insurance(msg.sender,insuranceIDGenerator,_name);
        insurance.push(h);
    }

    function func_deRegisterIC(address _icAddr) public
    {
        deRegister_IC[_icAddr]=true;
    }

    function func_securityMoneyIc(address _icAddr,uint _securityMoney) public
    {
        securityMoney_IC[_icAddr]=_securityMoney;
    }

    function getInsuranceCompanybyID(uint _insuranceID) view public returns(address payable,uint,string memory)
    {
        require(_insuranceID>=1 && _insuranceID<=insuranceIDGenerator);
        uint x = _insuranceID-1;
        return(insurance[x].icaddr,insurance[x].insuranceCompanyID,insurance[x].name);
    }

    function getInsuranceCompanybyAddress(address _insuranceAdd) view public returns(address payable,uint,string memory)
    {
        uint _insuranceID = mapFromAddToId_I[_insuranceAdd];
        require(_insuranceID>=1 && _insuranceID<=insuranceIDGenerator);
        uint x = _insuranceID-1;
        return(insurance[x].icaddr,insurance[x].insuranceCompanyID,insurance[x].name);
    }


    function addSecurityMoney(uint _amount) public payable
    {
    	uint _icID;
    	(,_icID,)=getInsuranceCompanybyAddress(msg.sender); // checking for valid IC ID..
    	securityMoney_IC[msg.sender] += _amount;
    }

    function withdrawnSecurityMoney(uint _amount) public
    {
    	uint _icID;
    	(,_icID,)=getInsuranceCompanybyAddress(msg.sender); // checking for valid IC ID..
    	require(_amount<=securityMoney_IC[msg.sender]);
    	securityMoney_IC[msg.sender] -= _amount;
    	msg.sender.transfer(_amount);
    }


    //Permission Data Structure...
    mapping (address=>mapping(address=>bool)) public read_permission_to_H;
    mapping (address=>mapping(address=>bool)) public write_permission_to_H;
    mapping (address=>mapping(address=>bool)) public read_permission_to_IC;

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
}
