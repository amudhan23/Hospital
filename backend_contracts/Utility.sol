pragma solidity >0.6.0 <=0.7.0;

contract Utility
{
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
