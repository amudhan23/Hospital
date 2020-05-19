package main

import(

		// "fmt"
		"os"
		"io"
		"log"
		"time"
		"context"
	 	"github.com/ethereum/go-ethereum/ethclient"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
		"github.com/ethereum/go-ethereum/crypto"
	  "github.com/ethereum/go-ethereum/common"
		"math/big"
		 import_contract_p_h "../compiledGoFiles_testNet/contract_p_h"
		 import_contract_p_dbo "../compiledGoFiles_testNet/contract_p_dbo"
		 import_contract_p_ic_dbo "../compiledGoFiles_testNet/contract_p_ic_dbo"
		 import_contract_rc_dbo "../compiledGoFiles_testNet/contract_rc_dbo"
		 import_registration "../compiledGoFiles_testNet/contract_registration"
)



const hashFunctionOutputBitSize=32

var con_address_registration common.Address
var con_address_P_HA_1 common.Address
var con_address_P_HA_2 common.Address
var con_address_P_DBO common.Address
var con_address_P_IC common.Address
var con_address_RC_DBO common.Address

type DBToResearcCommunity struct {
		AggregatedData []byte
}



func main(){


			logFilePath:="../log_testNet/log_registration_testNet.txt"

			var err error
			f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
					log.Fatal("error opening file: %v", err)
			}
			defer f.Close()
			wrt := io.MultiWriter(os.Stdout, f)
			log.SetOutput(wrt)



			client, err := ethclient.Dial("https://ropsten.infura.io/v3/c196cc55405844f69b80c2dbd7f734a8")
			if err != nil {
				log.Fatal(err)
			}

			ctx := context.Background()



			patient_PrivateKey,err:=crypto.HexToECDSA("Your PrivateKey")
			if err != nil {
				log.Fatal(err)
			}


			hospital_PrivateKey, err := crypto.HexToECDSA("Your PrivateKey")
			if err != nil {
				log.Fatal(err)
			}


			dataBaseOwner_PrivateKey, err := crypto.HexToECDSA("Your PrivateKey")
			if err != nil {
				log.Fatal(err)
			}

			insuranceCompany_PrivateKey, err := crypto.HexToECDSA("Your PrivateKey")
			if err != nil {
				log.Fatal(err)
			}


			govtAuthority_PrivateKey,err:=crypto.HexToECDSA("Your PrivateKey")
			if err != nil {
				log.Fatal(err)
			}


			researchCommunity_PrivateKey,err:=crypto.HexToECDSA("Your PrivateKey")
			if err != nil {
				log.Fatal(err)
			}



			patient:=bind.NewKeyedTransactor(patient_PrivateKey)
			hospital:=bind.NewKeyedTransactor(hospital_PrivateKey)
			govtAuthority:=bind.NewKeyedTransactor(govtAuthority_PrivateKey)
			dataBaseOwner:=bind.NewKeyedTransactor(dataBaseOwner_PrivateKey)
			insuranceCompany:=bind.NewKeyedTransactor(insuranceCompany_PrivateKey)
			researchCommunity:=bind.NewKeyedTransactor(researchCommunity_PrivateKey)



			govtAuthority_deploying_registration_start:=time.Now()

								_,tx,_,err:=import_registration.DeploySystemUsersInfo(govtAuthority,client)
								if err!=nil{
								    log.Fatal(err)
								}

								con_address_registration,err=bind.WaitDeployed(ctx,client,tx)
								if err!=nil{
									log.Fatal(err)
								}

			govtAuthority_deploying_registration_stop:=time.Now()
			prefix:=" party : govtAuthority "
			log.SetPrefix(prefix)
			log.Printf("Time govtAuthority_DeploySystemUsersInfo : %v \n\n", govtAuthority_deploying_registration_stop.Sub(govtAuthority_deploying_registration_start))
			log.SetPrefix(prefix)
			log.Println("Tx govtAuthority_DeploySystemUsersInfo : ",tx.Gas()," txHash : ",tx.Hash().Hex()," contractAddress_registration : ",con_address_registration.Hex(),"\n")




			govtAuthority_deploying_PHA2_start:=time.Now()

								_,tx,_,err=import_contract_p_h.DeploySCPHA2(govtAuthority,client)
								if err!=nil{
								    log.Fatal(err)
								}

								con_address_P_HA_2,err=bind.WaitDeployed(ctx,client,tx)
								if err!=nil{
									log.Fatal(err)
								}

			govtAuthority_deploying_PHA2_stop:=time.Now()
			prefix=" party : govtAuthority "
			log.SetPrefix(prefix)
			log.Printf("Time govtAuthority_DeployingPHA2 : %v \n\n", govtAuthority_deploying_PHA2_stop.Sub(govtAuthority_deploying_PHA2_start))
			log.SetPrefix(prefix)
			log.Println("Tx govtAuthority_DeployingPHA2 : ",tx.Gas()," txHash : ",tx.Hash().Hex()," contractAddress_P_HA_2 : ",con_address_P_HA_2.Hex(),"\n")




			govtAuthority_deploying_PHA1_start:=time.Now()

								_,tx,_,err=import_contract_p_h.DeploySCPHA1(govtAuthority,client,con_address_P_HA_2)
								if err!=nil{
								    log.Fatal(err)
								}

								con_address_P_HA_1,err=bind.WaitDeployed(ctx,client,tx)
								if err!=nil{
									log.Fatal(err)
								}

			govtAuthority_deploying_PHA1_stop:=time.Now()
			prefix=" party : govtAuthority "
			log.SetPrefix(prefix)
			log.Printf("Time govtAuthority_DeployingPHA1 : %v \n\n", govtAuthority_deploying_PHA1_stop.Sub(govtAuthority_deploying_PHA1_start))
			log.SetPrefix(prefix)
			log.Println("Tx govtAuthority_DeployingPHA1 : ",tx.Gas()," txHash : ",tx.Hash().Hex()," contractAddress_P_HA_1 : ",con_address_P_HA_1.Hex(),"\n")


			// instance_PHA1,err:=import_contract_p_h.NewSCPHA1(con_address_P_HA_1,client)
			// if err!=nil{
			// 		log.Fatal(err)
			// }
			//
			//
			//
			// AddrSCPHA2,err:=instance_PHA1.AddrSCPHA2(nil)
			// if err!=nil{
			// 		log.Fatal(err)
			// }
			//
			//
			// fmt.Println("\n\nAddrSCPHA2 : ",AddrSCPHA2.Hex())

			// con_address_P_HA_2=AddrSCPHA2



			govtAuthority_deploying_PDBO_start:=time.Now()


									_,tx,_,err=import_contract_p_dbo.DeploySCPDBO(govtAuthority,client)
									if err!=nil{
											log.Fatal(err)
									}

									con_address_P_DBO,err=bind.WaitDeployed(ctx,client,tx)
									if err!=nil{
										log.Fatal(err)
									}



			govtAuthority_deploying_PDBO_stop:=time.Now()
			prefix=" party : govtAuthority "
			log.SetPrefix(prefix)
			log.Printf("Time govtAuthority_DeployingPDBO : %v \n\n", govtAuthority_deploying_PDBO_stop.Sub(govtAuthority_deploying_PDBO_start))
			log.SetPrefix(prefix)
			log.Println("Tx govtAuthority_DeployingPDBO : ",tx.Gas()," txHash : ",tx.Hash().Hex()," con_address_P_DBO : ",con_address_P_DBO.Hex(),"\n")



			govtAuthority_deploying_PIC_start:=time.Now()


									_,tx,_,err=import_contract_p_ic_dbo.DeploySCPICDBO(govtAuthority,client)
									if err!=nil{
											log.Fatal(err)
									}

									con_address_P_IC,err=bind.WaitDeployed(ctx,client,tx)
									if err!=nil{
										log.Fatal(err)
									}


			govtAuthority_deploying_PIC_stop:=time.Now()
			prefix=" party : govtAuthority "
			log.SetPrefix(prefix)
			log.Printf("Time govtAuthority_DeployingPIC : %v \n\n", govtAuthority_deploying_PIC_stop.Sub(govtAuthority_deploying_PIC_start))
			log.SetPrefix(prefix)
			log.Println("Tx govtAuthority_DeployingPIC : ",tx.Gas()," txHash : ",tx.Hash().Hex()," con_address_P_IC : ",con_address_P_IC.Hex(),"\n")


			govtAuthority_deployingSCRCDBO_start:=time.Now()

								_,tx,_,err=import_contract_rc_dbo.DeploySCRCDBO(govtAuthority,client)
								if err!=nil{
								    log.Fatal(err)
								}

								con_address_RC_DBO,err=bind.WaitDeployed(ctx,client,tx)
								if err!=nil{
									log.Fatal(err)
								}

			govtAuthority_deployingSCRCDBO_stop:=time.Now()
			prefix=" party : GovtAuthority "
			log.SetPrefix(prefix)
			log.Printf("Time GovtAuthority_DeployingSCRCDBO : %v \n\n", govtAuthority_deployingSCRCDBO_stop.Sub(govtAuthority_deployingSCRCDBO_start))
			log.SetPrefix(prefix)
			log.Println("Tx GovtAuthority_DeploySCRCDBO : ",tx.Gas()," txHash : ",tx.Hash().Hex()," con_address_RC_DBO : ",con_address_RC_DBO.Hex(),"\n")



			instance_registration,err:=import_contract_p_h.NewSystemUsersInfo(con_address_registration,client)
			if err!=nil{
					log.Fatal(err)
			}



			patient_registration_start:=time.Now()


							patient_Name:="Patient1"
							patient_Age:=20
							patient_mobile:=8123456789
							patient_addr:="Area-1,Chennai"


							tx,err=instance_registration.PatientRegistration(&bind.TransactOpts{
													From:patient.From,
													Signer:patient.Signer,
							},patient_Name,big.NewInt(int64(patient_Age)),big.NewInt(int64(patient_mobile)),patient_addr)
							if err!=nil{
					    		log.Fatal(err)
							}

							_,err=bind.WaitMined(ctx,client,tx)
							if err!=nil{
								log.Fatal(err)
							}


			patient_registration_stop:=time.Now()
			prefix=" party : Patient "
			log.SetPrefix(prefix)
			log.Printf("Time Patient_Registration : %v \n\n", patient_registration_stop.Sub(patient_registration_start))
			log.SetPrefix(prefix)
			log.Println("Tx Patient_Registration : ",tx.Gas()," txHash : ",tx.Hash().Hex(),"\n")



			hospital_registration_start:=time.Now()


						 hospital_Name:="Hospital1"

							tx,err=instance_registration.HospitalRegistration(&bind.TransactOpts{
													From:hospital.From,
													Signer:hospital.Signer,
							},hospital_Name)
							if err!=nil{
					    		log.Fatal(err)
							}

							_,err=bind.WaitMined(ctx,client,tx)
							if err!=nil{
								log.Fatal(err)
							}


			hospital_registration_stop:=time.Now()
			prefix=" party : Hospital "
			log.SetPrefix(prefix)
			log.Printf("Time Hospital_Registration : %v \n\n", hospital_registration_stop.Sub(hospital_registration_start))
			log.SetPrefix(prefix)
			log.Println("Tx Hospital_Registration : ",tx.Gas()," txHash : ",tx.Hash().Hex(),"\n")



			dataBaseOwner_registration_start:=time.Now()


							dboName:="dbo1"


							tx,err=instance_registration.DboRegistration(&bind.TransactOpts{
													From:govtAuthority.From,
													Signer:govtAuthority.Signer,
							},dataBaseOwner.From,dboName)
							if err!=nil{
					    		log.Fatal(err)
							}

							_,err=bind.WaitMined(ctx,client,tx)
							if err!=nil{
								log.Fatal(err)
							}


			dataBaseOwner_registration_stop:=time.Now()
			prefix=" party : GovtAuthority "
			log.SetPrefix(prefix)
			log.Printf("Time DataBaseOwner_Registration : %v \n\n", dataBaseOwner_registration_stop.Sub(dataBaseOwner_registration_start))
			log.SetPrefix(prefix)
			log.Println("Tx DataBaseOwner_Registration : ",tx.Gas()," txHash : ",tx.Hash().Hex(),"\n")



			insuranceCompany_registration_start:=time.Now()


						 insuranceCompany_Name:="insuranceCompany1"

							tx,err=instance_registration.InsuranceCompanyRegistration(&bind.TransactOpts{
													From:govtAuthority.From,
													Signer:govtAuthority.Signer,
							},insuranceCompany.From,insuranceCompany_Name)
							if err!=nil{
					    		log.Fatal(err)
							}


							_,err=bind.WaitMined(ctx,client,tx)
							if err!=nil{
								log.Fatal(err)
							}


			insuranceCompany_registration_stop:=time.Now()
			prefix=" party : GovtAuthority "
			log.SetPrefix(prefix)
			log.Printf("Time InsuranceCompany_Registration : %v \n\n", insuranceCompany_registration_stop.Sub(insuranceCompany_registration_start))
			log.SetPrefix(prefix)
			log.Println("Tx InsuranceCompany_Registration : ",tx.Gas()," txHash : ",tx.Hash().Hex(),"\n")



			researchCommunity_registration_start:=time.Now()


						 researchCommunityName:="rc1"

							tx,err=instance_registration.RcRegistration(&bind.TransactOpts{
													From:govtAuthority.From,
													Signer:govtAuthority.Signer,
							},researchCommunity.From,researchCommunityName)
							if err!=nil{
					    		log.Fatal(err)
							}


							_,err=bind.WaitMined(ctx,client,tx)
							if err!=nil{
								log.Fatal(err)
							}


			researchCommunity_registration_stop:=time.Now()
			prefix=" party : GovtAuthority "
			log.SetPrefix(prefix)
			log.Printf("Time ResearchCommunity_registration : %v \n\n", researchCommunity_registration_stop.Sub(researchCommunity_registration_start))
			log.SetPrefix(prefix)
			log.Println("Tx ResearchCommunity_registration : ",tx.Gas()," txHash : ",tx.Hash().Hex(),"\n")


	}
