package main

import(

		"fmt"
		 kec "github.com/ethereum/go-ethereum/crypto"
		"os"
		"io"
		"log"
		"encoding/json"
		"time"
		"sync"
		"reflect"
		"strconv"
		"crypto/rand"
		"crypto/rsa"
	  "crypto/ecdsa"
	 	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
		"github.com/ethereum/go-ethereum/crypto"
		"github.com/ethereum/go-ethereum/core"
	  "github.com/ethereum/go-ethereum/common"
		"github.com/ethereum/go-ethereum/core/types"
		"crypto/sha256"
		"math/big"
		"crypto/x509"
    "encoding/pem"
		 mathlib "math"
	   import_contract_p_dbo "../compiledGoFiles_backend/contract_p_dbo"
		 import_contract_p_ic_dbo "../compiledGoFiles_backend/contract_p_ic_dbo"
		 import_contract_registration "../compiledGoFiles_backend/contract_registration"

)


var buffer_size uint
var noOfInputGates uint
var totalNumberOfGates uint


var con_address_registration common.Address
var con_address_P_HA_1 common.Address
var con_address_P_HA_2 common.Address
var con_address_P_DO common.Address
var con_address_P_IC_DBO common.Address


const keySize=32
const maxLinesToGate =2
const hashFunctionOutputBitSize=32


const timeLimitForPatientLocking=100
const timeLimitForStartingOperation=100
const timeLimitForVerifyingSignature=100
const timeLimitForGeneratingFillBill=100
const timeLimitForSigningFinalBill=100
const timeLimitForKeyRevealing=100
const timeLimitForAcceptingOrProducingComplain=100



const timeLimitForMerkleRootVerification_P_DO=100
const timeLimitForKeyReveal_P_DO=100
const timeLimitForGivingComplainOrPuttingInTheDatatbase_P_DO=100



const timeLimiitForPolicyBuying=100
const timeLimitForICLocking=100
const timeLimitForKeyReveal=100
const timeLimitForFinalApproval=100


type HospitalToPatient struct {
		EncryptedOutputOfGates [][]byte
		HospitalAddress common.Address
}

type PatientToDataBaseOwner struct {
		EncryptedOutputOfGates [][]byte
}

type PatientICToDataBaseOwner struct {
		InsuranceComapanyId int
		EstimatedBillID int
		ClaimId int
}

type DataBaseOwnerToInsuranceCompany struct {
		EncryptedOutputOfGates [][]byte
}


func main(){


			noOfInputGates=4
			buffer_size=32
			totalNumberOfGates=2*noOfInputGates

			logFilePath:="../log_backend/log_p_ic_dbo_backend.txt"

			var err error
			f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
					// log.Fatalf("error opening file: %v", err)
			defer f.Close()
			}
			wrt := io.MultiWriter(os.Stdout, f)
			log.SetOutput(wrt)



			alloc := make(core.GenesisAlloc)

			patient,_,_:=	AuthAndAddressGeneration()
			hospital,_,hospitalPrivateKey:= AuthAndAddressGeneration()
			dataBaseOwner,_,dbprivateKey:= AuthAndAddressGeneration()
			govtAuthority,_,_:=AuthAndAddressGeneration()
			insurance_company,_,_:=AuthAndAddressGeneration()
			alloc[patient.From] = core.GenesisAccount{Balance: big.NewInt(1100000000000000)}
			alloc[hospital.From] = core.GenesisAccount{Balance: big.NewInt(11000000000)}
			alloc[dataBaseOwner.From] = core.GenesisAccount{Balance: big.NewInt(11000000000)}
			alloc[insurance_company.From] = core.GenesisAccount{Balance: big.NewInt(11000000000000)}
			alloc[govtAuthority.From] = core.GenesisAccount{Balance: big.NewInt(11000000000000)}

			client := backends.NewSimulatedBackend(alloc,6000000000000)


			con_address_registration,_,_,err=import_contract_p_dbo.DeploySystemUsersInfo(govtAuthority,client)
			if err!=nil{
			    log.Fatal(err)
			}
			client.Commit()


			con_address_P_HA_2,_,_,err=import_contract_p_dbo.DeploySCPHA2(govtAuthority,client)
			if err!=nil{
			    log.Fatal(err)
			  }
			client.Commit()



			con_address_P_HA_1,_,_,err=import_contract_p_dbo.DeploySCPHA1(govtAuthority,client,con_address_P_HA_2)
			if err!=nil{
			    log.Fatal(err)
			 }
			client.Commit()





			con_address_P_DO,_,_,err=import_contract_p_dbo.DeploySCPDBO(govtAuthority,client)
			if err!=nil{
			    log.Fatal(err)
			}
			client.Commit()


			con_address_P_IC_DBO,_,_,err=import_contract_p_ic_dbo.DeploySCPICDBO(govtAuthority,client)
			if err!=nil{
			    log.Fatal(err)
			}
			client.Commit()



			var wg sync.WaitGroup

			channel_HospitalToPatient_EstimatedBillId:=make(chan int)
			channel_HospitalToPatient_EncryptedFile_Object:=make(chan []byte)
			channel_PatientToHospital_PatientId:=make(chan int)
			channel_HospitalToPatient_HospitalId:=make(chan int)
			channel_HospitalToPatient_sendHashDateMRenc:=make(chan [hashFunctionOutputBitSize]byte)
			channel_multiSignId:=make(chan int)
			// channel_HospitalToPatient_send_HashOf_PatientId_DateOfReport_HashOfEncryptedFile:=make(chan [hashFunctionOutputBitSize]byte)
			// channel_HospitalToPatient_HospitalId:=make(chan [hashFunctionOutputBitSize])

			wg.Add(2)

			go Patient(patient,client,channel_HospitalToPatient_EstimatedBillId,channel_HospitalToPatient_EncryptedFile_Object,channel_PatientToHospital_PatientId,channel_HospitalToPatient_HospitalId,channel_multiSignId,&wg)

			go Hospital(hospital,client,channel_HospitalToPatient_EncryptedFile_Object,channel_HospitalToPatient_EstimatedBillId,channel_PatientToHospital_PatientId,channel_HospitalToPatient_HospitalId,channel_HospitalToPatient_sendHashDateMRenc,channel_multiSignId,&wg,hospitalPrivateKey)

			wg.Wait()


			channel_Patient_DO_DBOid:=make(chan int)
			channel_Patient_DO_appId:=make(chan int)
			channel_Patient_DO_EncryptedFile_Object:=make(chan []byte)


			var wg2 sync.WaitGroup
			wg2.Add(2)

			instance_registration,err								:=import_contract_registration.NewSystemUsersInfo(con_address_registration,client)


			govtAuthority_dboRegistration_start:=time.Now()



								dboName:="dbo1"

								tx,err:=instance_registration.DboRegistration(&bind.TransactOpts{
										From:govtAuthority.From,
										Signer:govtAuthority.Signer,
								},dataBaseOwner.From,dboName)
								if err!=nil{
										log.Fatal(err)
								}
								client.Commit()



			govtAuthority_dboRegistration_stop:=time.Now()
			prefix:=""
			log.SetPrefix(prefix)
			log.Printf("Time GovtAuthority_DBO_Registration : %v \n\n", govtAuthority_dboRegistration_stop.Sub(govtAuthority_dboRegistration_start))
			log.SetPrefix(prefix)
			log.Println("Tx GovtAuthority_DBO_Registration : ",tx.Gas(),"\n")


			go DatabaseOwner(dataBaseOwner,client,channel_Patient_DO_appId,channel_Patient_DO_EncryptedFile_Object,channel_Patient_DO_DBOid,&wg2)


			go Patient_Do(patient,client,channel_Patient_DO_appId,channel_Patient_DO_EncryptedFile_Object,channel_Patient_DO_DBOid,&wg2)


			wg2.Wait()


			var wg3 sync.WaitGroup
			wg3.Add(3)

			channel_PatientIC_IC_patientId:=make(chan int)
			channel_PatientIC_IC_insuranceCompanyId:=make(chan int)
			channel_PatientIC_IC_claimId:=make(chan int)
			channel_PatientIC_DO_SendIcDetails:=make(chan []byte)
			channel_DOIC_insuranceCompany_SendEncryptedFile:=make(chan []byte)


			govtAuthority_icRegistration_start:=time.Now()



								icName:="ic1"

								tx,err=instance_registration.InsuranceCompanyRegistration(&bind.TransactOpts{
										From:govtAuthority.From,
										Signer:govtAuthority.Signer,
								},insurance_company.From,icName)
								if err!=nil{
										log.Fatal(err)
								}
								client.Commit()



			govtAuthority_icRegistration_stop:=time.Now()
			prefix=""
			log.SetPrefix(prefix)
			log.Printf("Time GovtAuthority_IC_Registration : %v \n\n", govtAuthority_icRegistration_stop.Sub(govtAuthority_icRegistration_start))
			log.SetPrefix(prefix)
			log.Println("Tx GovtAuthority_IC_Registration : ",tx.Gas(),"\n")



			go Patient_IC(patient,client,channel_PatientIC_IC_patientId,channel_PatientIC_IC_insuranceCompanyId,channel_PatientIC_IC_claimId,channel_PatientIC_DO_SendIcDetails,dataBaseOwner.From,&wg3)

			go Insurance_company(insurance_company,client,channel_PatientIC_IC_patientId,channel_PatientIC_IC_insuranceCompanyId,channel_PatientIC_IC_claimId,channel_DOIC_insuranceCompany_SendEncryptedFile,&wg3)

			go DataBaseOwner_InsuranceCompany(dataBaseOwner,client,channel_PatientIC_DO_SendIcDetails,channel_DOIC_insuranceCompany_SendEncryptedFile,dbprivateKey,&wg3)

			wg3.Wait()




	}

	func DataBaseOwner_InsuranceCompany(dataBaseOwner *bind.TransactOpts,client *backends.SimulatedBackend,channel_PatientIC_DO_SendIcDetails chan []byte,channel_DOIC_insuranceCompany_SendEncryptedFile chan []byte,dbprivateKey *ecdsa.PrivateKey,wg *sync.WaitGroup){


				prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : DataBaseOwner_InsuranceCompany "



				instance_P_IC_DBO,err:=import_contract_p_ic_dbo.NewSCPICDBO(con_address_P_IC_DBO,client)

				var PatientICToDataBaseOwner_object PatientICToDataBaseOwner
				var DataBaseOwnerToInsuranceCompany_object DataBaseOwnerToInsuranceCompany

				byteArray_PatientICToDataBaseOwner:=<-channel_PatientIC_DO_SendIcDetails


				dbIC_computing_signature_start:=time.Now()


									err=json.Unmarshal(byteArray_PatientICToDataBaseOwner, &PatientICToDataBaseOwner_object)
									if err != nil {
											log.Fatal(err)
									}


									claimId:=PatientICToDataBaseOwner_object.ClaimId


									claimDetails_object,err:=instance_P_IC_DBO.ClaimDetails(nil,big.NewInt(int64(claimId-1)))
									if err!=nil{
											log.Fatal(err)
									}



									rsa_publicKey_bytes:=claimDetails_object.PublicKey
									rsa_publicKey:=BytesToPublicKey(rsa_publicKey_bytes)



									inputVectors:=FileRead()
									encryptedFileChunks:=EncryptFileChunks(inputVectors,rsa_publicKey)
									MR_encryptedFileChunks:=createMerkleRootForInputVectors(encryptedFileChunks)


									byte_MR_encryptedFileChunks:=make([]byte,32)
									for i:=0;i<32;i++{
											byte_MR_encryptedFileChunks[i]=MR_encryptedFileChunks[i]
									}


								signature_db_MR_encryptedFileChunks,err := crypto.Sign(byte_MR_encryptedFileChunks,dbprivateKey)
								if err != nil {
									log.Fatal(err)
								}



				dbIC_computing_signature_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time DbIC_computing_signature : %v \n\n", dbIC_computing_signature_stop.Sub(dbIC_computing_signature_start))



				dbIC_KeepSignature_start:=time.Now()

								tx,err:=instance_P_IC_DBO.DboKeepSigOnHashOfEncFile(&bind.TransactOpts{
										From:dataBaseOwner.From,
										Signer:dataBaseOwner.Signer,
								},big.NewInt(int64(claimId)),signature_db_MR_encryptedFileChunks)
								if err!=nil{
										log.Fatal(err)
								}
								client.Commit()


				dbIC_KeepSignature_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time DbIC_KeepSignature : %v \n\n", dbIC_KeepSignature_stop.Sub(dbIC_KeepSignature_start))
				log.SetPrefix(prefix)
				log.Println("Tx DbIC_KeepSigOnHashOfEncFile : ",tx.Gas(),"\n")




				dbIC_sendingEcncryptedFilesToIC_start:=time.Now()


									DataBaseOwnerToInsuranceCompany_object.EncryptedOutputOfGates=encryptedFileChunks
									byteArrayFromDataBaseOwnerToInsuranceCompany,err := json.Marshal(DataBaseOwnerToInsuranceCompany_object)
									if err != nil {
												log.Fatal(err)
									}

									channel_DOIC_insuranceCompany_SendEncryptedFile<-byteArrayFromDataBaseOwnerToInsuranceCompany



				dbIC_sendingEcncryptedFilesToIC_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time DbIC_sendingEcncryptedFilesToIC : %v \n\n", dbIC_sendingEcncryptedFilesToIC_stop.Sub(dbIC_sendingEcncryptedFilesToIC_start))


				wg.Done()




	}



	func Patient_IC(patient_IC *bind.TransactOpts,client *backends.SimulatedBackend,channel_PatientIC_IC_patientId chan int,channel_PatientIC_IC_insuranceCompanyId chan int,channel_PatientIC_IC_claimId chan int,channel_PatientIC_DO_SendIcDetails chan []byte,dataBaseOwner_Address common.Address,wg *sync.WaitGroup) {


			prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : Patient_WithInsuranceCompany "

			instance_P_IC_DBO,err:=import_contract_p_ic_dbo.NewSCPICDBO(con_address_P_IC_DBO,client)
			// instance_P_DBO,err:=import_contract_p_dbo.NewSCPDBO(con_address_P_DO,client)

			instance_registration,err:=import_contract_p_dbo.NewSystemUsersInfo(con_address_registration,client)
			if err!=nil{
			    log.Fatal(err)
			}

			var tx *types.Transaction
			var	PatientICToDataBaseOwner_object PatientICToDataBaseOwner

			patientId:=1
			insuranceCompanyId:=<-channel_PatientIC_IC_insuranceCompanyId


			patientIC_buyPolicyPhaseOne_start:=time.Now()

						policyPrice:=50
						claimAmount:=30
						estimatedBillId:=1
						var insuranceTermsAndConditionsFileHash [32]byte

						tx,err=instance_P_IC_DBO.BuyPolicyPhaseOne(&bind.TransactOpts{
								From:patient_IC.From,
								Signer:patient_IC.Signer,
								Value:big.NewInt(int64(policyPrice)),
						},con_address_registration,big.NewInt(int64(insuranceCompanyId)),insuranceTermsAndConditionsFileHash)
						if err!=nil{
								log.Fatal(err)
						}
						client.Commit()

			patientIC_buyPolicyPhaseOne_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time PatientIC_buyPolicyPhaseOne : %v \n\n", patientIC_buyPolicyPhaseOne_stop.Sub(patientIC_buyPolicyPhaseOne_start))
			log.SetPrefix(prefix)
			log.Println("Tx PatientIC_buyPolicyPhaseOne : ",tx.Gas(),"\n")


			channel_PatientIC_IC_patientId<-patientId

			var PolicyDetailId *big.Int

			patientIC_WaitingForICtoLockMoney_start:=time.Now()

								patient_IcLockMoney:=false

								insuranceCompanyAddr,_,_,err:=instance_registration.GetInsuranceCompanybyID(nil,big.NewInt(int64(insuranceCompanyId)))
								if err!=nil{
										log.Fatal(err)
								}


								for i:=0;i<timeLimiitForPolicyBuying;i++ {


											PolicyDetailId,err=instance_P_IC_DBO.PolicyIdMap(nil,patient_IC.From,insuranceCompanyAddr)
											if err!=nil{
													log.Fatal(err)
											}

											if PolicyDetailId.Int64()!=0	{
												patient_IcLockMoney=true
												break
											}
											time.Sleep(time.Second *1)

							 }


			patientIC_WaitingForICtoLockMoney_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time PatientIC_WaitingForICtoLockMoney : %v \n\n", patientIC_WaitingForICtoLockMoney_stop.Sub(patientIC_WaitingForICtoLockMoney_start))



			if !patient_IcLockMoney{

					 		patientIC_UnlockingAmount_ICnotLockingMoney_start:=time.Now()


								 				PolicyPhaseOne_object,err:=instance_P_IC_DBO.PolicyPhaseOne(nil,patient_IC.From,insuranceCompanyAddr)
												if err!=nil{
														log.Fatal(err)
												}


												Time:=PolicyPhaseOne_object.Time

												Now,err:=instance_P_IC_DBO.Now(nil)
												if err!=nil{
													    log.Fatal(err)
												}



												for {


																	if Now.Int64()>Time.Int64()+ timeLimiitForPolicyBuying{
																		break
																	}

																	_,err:=instance_P_IC_DBO.IncreaseTheTimeBy10seconds(&bind.TransactOpts{
																				From:patient_IC.From,
																				Signer:patient_IC.Signer,
																	})
																	if err!=nil{
																			log.Fatal(err)
																	}
																	client.Commit()

																	Now,err=instance_P_IC_DBO.Now(nil)
																	if err!=nil{
																			log.Fatal(err)
																	}



												}


												tx,err=instance_P_IC_DBO.PatientWithdrawLockedPolicyBuyMoney(&bind.TransactOpts{
															From:patient_IC.From,
															Signer:patient_IC.Signer,
												},con_address_registration,big.NewInt(int64(insuranceCompanyId)))
												if err!=nil{
												    log.Fatal(err)
												}
												client.Commit()



						 patientIC_UnlockingAmount_ICnotLockingMoney_stop:=time.Now()
						 log.SetPrefix(prefix)
			 			 log.Printf("Time PatientIC_UnlockingAmount_ICnotLockingMoney : %v \n\n", patientIC_UnlockingAmount_ICnotLockingMoney_stop.Sub(patientIC_UnlockingAmount_ICnotLockingMoney_start))
						 log.SetPrefix(prefix)
			 			 log.Println("Tx PatientIC_PatientWithdrawLockedPolicyBuyMoney : ",tx.Gas(),"\n")

						 wg.Done()


				}


				patientIC_GeneratingClaim_start:=time.Now()



								appId:=1


								rsa_privateKey,rsa_publicKey:=GenerateKeyPair(int(buffer_size*32))

								bytes_rsa_privateKey:=PrivateKeyToBytes(rsa_privateKey)
								bytes_rsa_publicKey:=PublicKeyToBytes(rsa_publicKey)

								rsa_privateKey_commit:=rsa_fnKeycommit(bytes_rsa_privateKey)


								tx,err=instance_P_IC_DBO.ClaimMoney(&bind.TransactOpts{
											From:patient_IC.From,
											Signer:patient_IC.Signer,
								},con_address_registration,con_address_P_HA_1,con_address_P_DO,PolicyDetailId,big.NewInt(int64(claimAmount)),big.NewInt(int64(estimatedBillId)),bytes_rsa_publicKey,rsa_privateKey_commit,dataBaseOwner_Address,big.NewInt(int64(appId)))
								if err!=nil{
										log.Fatal(err)
								}
								client.Commit()



				patientIC_GeneratingClaim_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientIC_GeneratingClaim : %v \n\n", patientIC_GeneratingClaim_stop.Sub(patientIC_GeneratingClaim_start))
				log.SetPrefix(prefix)
				log.Println("Tx PatientIC_GenerateClaim : ",tx.Gas(),"\n")




				patientIC_sendClaimIdToIC_start:=time.Now()


								PolicyDetailId_int:=PolicyDetailId.Int64()
								claims,err:=instance_P_IC_DBO.ReturnClaimIds(nil,big.NewInt(int64(PolicyDetailId_int)))
								claimId:=claims[len(claims)-1].Int64()
								channel_PatientIC_IC_claimId<-int(claimId)


				patientIC_sendClaimIdToIC_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientIC_SendClaimIdToIC : %v \n\n", patientIC_sendClaimIdToIC_stop.Sub(patientIC_sendClaimIdToIC_start))


				patientIC_sendClaimIdToDataBaseOwner_start:=time.Now()

								PatientICToDataBaseOwner_object.InsuranceComapanyId=insuranceCompanyId
								PatientICToDataBaseOwner_object.EstimatedBillID=estimatedBillId
								PatientICToDataBaseOwner_object.ClaimId=int(claimId)

								byteArrayFromPatientToDateBaseOwner,err := json.Marshal(PatientICToDataBaseOwner_object)
								if err != nil {
											log.Fatal(err)
								}

								channel_PatientIC_DO_SendIcDetails<-byteArrayFromPatientToDateBaseOwner

				patientIC_sendClaimIdToDataBaseOwner_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientIC_sendClaimIdToDataBaseOwner : %v \n\n", patientIC_sendClaimIdToDataBaseOwner_stop.Sub(patientIC_sendClaimIdToDataBaseOwner_start))




				patientIC_waitingForInsuranceCompanyToApprovedAmount_start:=time.Now()

								patientIC_insuranceCompanyApprovedClaim:=false


								for i:=0;i<timeLimitForICLocking;i++ {


											claimDetails_object,err:=instance_P_IC_DBO.ClaimDetails(nil,big.NewInt(int64(claimId-1)))
											if err!=nil{
													log.Fatal(err)
											}

											TimestampIcLocking:=claimDetails_object.TimestampIcLocking


											if TimestampIcLocking.Int64()!=0	{
												patientIC_insuranceCompanyApprovedClaim=true
												break
											}
											time.Sleep(time.Second *1)

							 }


				patientIC_waitingForInsuranceCompanyToApprovedAmount_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientIC_waitingForInsuranceCompanyToApprovedAmount : %v \n\n", patientIC_waitingForInsuranceCompanyToApprovedAmount_stop.Sub(patientIC_waitingForInsuranceCompanyToApprovedAmount_start))


				if !patientIC_insuranceCompanyApprovedClaim{

							wg.Done()

				}


				patientIC_revealSecretKey_start:=time.Now()


									tx,err=instance_P_IC_DBO.RevealSecretKey(&bind.TransactOpts{
												From:patient_IC.From,
												Signer:patient_IC.Signer,
									},con_address_registration,bytes_rsa_privateKey,big.NewInt(int64(claimId)),big.NewInt(int64(insuranceCompanyId)))
									if err!=nil{
											log.Fatal(err)
									}
									client.Commit()


				patientIC_revealSecretKey_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientIC_RevealKey : %v \n\n", patientIC_revealSecretKey_stop.Sub(patientIC_revealSecretKey_start))
				log.SetPrefix(prefix)
				log.Println("Tx PatientIC_RevealSecretKey : ",tx.Gas(),"\n")



				patientIC_waitingForICtoApprove_start:=time.Now()

								patientIC_ICapprove:=false

								for i:=0;i<timeLimitForFinalApproval;i++ {


											claimDetails_object,err:=instance_P_IC_DBO.ClaimDetails(nil,big.NewInt(int64(claimId-1)))
											if err!=nil{
													log.Fatal(err)
											}

											TimestampApproval:=claimDetails_object.TimestampApproval

											if TimestampApproval.Int64()!=0	{
												patientIC_ICapprove=true
												break
											}
											time.Sleep(time.Second *1)

							 }



				patientIC_waitingForICtoApprove_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientIC_waitingForICtoApprove : %v \n\n", patientIC_waitingForICtoApprove_stop.Sub(patientIC_waitingForICtoApprove_start))



				if !patientIC_ICapprove{

										patientIC_selfApproval_start:=time.Now()


															claimDetails_object,err:=instance_P_IC_DBO.ClaimDetails(nil,big.NewInt(int64(claimId-1)))
															if err!=nil{
																	log.Fatal(err)
															}



															TimestampKeyReveal:=claimDetails_object.TimestampKeyReveal

															Now,err:=instance_P_IC_DBO.Now(nil)
															if err!=nil{
																		log.Fatal(err)
															}



															for {


																				if Now.Int64()>TimestampKeyReveal.Int64()+ timeLimitForFinalApproval{
																					break
																				}

																				_,err:=instance_P_IC_DBO.IncreaseTheTimeBy10seconds(&bind.TransactOpts{
																							From:patient_IC.From,
																							Signer:patient_IC.Signer,
																				})
																				if err!=nil{
																						log.Fatal(err)
																				}
																				client.Commit()

																				Now,err=instance_P_IC_DBO.Now(nil)
																				if err!=nil{
																						log.Fatal(err)
																				}



															}

															tx,err=instance_P_IC_DBO.SelfApproveClaim(&bind.TransactOpts{
																		From:patient_IC.From,
																		Signer:patient_IC.Signer,
															},con_address_registration,big.NewInt(int64(claimId)),big.NewInt(int64(insuranceCompanyId)))
															if err!=nil{
																	log.Fatal(err)
															}
															client.Commit()



							patientIC_selfApproval_stop:=time.Now()
							log.SetPrefix(prefix)
							log.Printf("Time PatientIC_selfApproval : %v \n\n", patientIC_selfApproval_stop.Sub(patientIC_selfApproval_start))
							log.SetPrefix(prefix)
							log.Println("Tx PatientIC_SelfApproval : ",tx.Gas(),"\n")
							wg.Done()

					}

					wg.Done()

	}


	func Insurance_company(insuranceCompany *bind.TransactOpts,client *backends.SimulatedBackend,channel_PatientIC_IC_patientId chan int,channel_PatientIC_IC_insuranceCompanyId chan int,channel_PatientIC_IC_claimId chan int,channel_DOIC_insuranceCompany_SendEncryptedFile chan []byte,wg *sync.WaitGroup){

			prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : InsuranceCompany "


			var tx *types.Transaction
			var DataBaseOwnerToInsuranceCompany_object DataBaseOwnerToInsuranceCompany


			instance_registration,err:=import_contract_p_dbo.NewSystemUsersInfo(con_address_registration,client)
			if err!=nil{
			    log.Fatal(err)
			}

			instance_P_HA_1,err:=import_contract_p_dbo.NewSCPHA1(con_address_P_HA_1,client)
			if err!=nil{
			    log.Fatal(err)
			}

			_,insuranceCompanyId,_,err:=instance_registration.GetInsuranceCompanybyAddress(nil,insuranceCompany.From)
			if err!=nil{
					log.Fatal(err)
			}

			fmt.Println("\ninsuranceCompanyId : ",insuranceCompanyId,"\n")

			channel_PatientIC_IC_insuranceCompanyId<-int(insuranceCompanyId.Int64())


			instance_P_IC_DBO,err:=import_contract_p_ic_dbo.NewSCPICDBO(con_address_P_IC_DBO,client)




			insuranceCompany_depositSecurityMoney_start:=time.Now()

							securityMoney:=100

							tx,err=instance_P_IC_DBO.AddSecurityMoney(&bind.TransactOpts{
									From:insuranceCompany.From,
									Signer:insuranceCompany.Signer,
									Value:big.NewInt(int64(securityMoney)),
							},con_address_registration,big.NewInt(int64(securityMoney)))
							if err!=nil{
									log.Fatal(err)
							}
							client.Commit()



			insuranceCompany_depositSecurityMoney_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time InsuranceCompany_SecurityDeposist : %v \n\n", insuranceCompany_depositSecurityMoney_stop.Sub(insuranceCompany_depositSecurityMoney_start))
			log.SetPrefix(prefix)
			log.Println("Tx InsuranceCompany_SecurityDeposit : ",tx.Gas(),"\n")









			patientId:=<-channel_PatientIC_IC_patientId


			insuranceCompany_LockPolicyPhaseTwo_start:=time.Now()

							policyPrice:=50;

							var insuranceTermsAndConditionsFileHash [32]byte


							tx,err=instance_P_IC_DBO.BuyPolicyPhaseTwo(&bind.TransactOpts{
									From:insuranceCompany.From,
									Signer:insuranceCompany.Signer,
									// Value:big.NewInt(int64(5000000)),
							},con_address_registration,big.NewInt(int64(patientId)),big.NewInt(int64(policyPrice)),insuranceTermsAndConditionsFileHash)
							if err!=nil{
									log.Fatal(err)
							}
							client.Commit()


			insuranceCompany_LockPolicyPhaseTwo_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time InsuranceCompany_LockPolicyPhaseTwo : %v \n\n", insuranceCompany_LockPolicyPhaseTwo_stop.Sub(insuranceCompany_LockPolicyPhaseTwo_start))
			log.SetPrefix(prefix)
			log.Println("Tx InsuranceCompany_LockPolicyPhaseTwo : ",tx.Gas(),"\n")


			insuranceCompany_waitingForPatientToSendClaimId_start:=time.Now()

								claimId:=<-channel_PatientIC_IC_claimId


			insuranceCompany_waitingForPatientToSendClaimId_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time InsuranceCompany_WaitingForPatientToSendClaimId : %v \n\n", insuranceCompany_waitingForPatientToSendClaimId_stop.Sub(insuranceCompany_waitingForPatientToSendClaimId_start))


			insuranceCompany_GettingEncFileAndVerifyingDbSignature_start:=time.Now()

								byteArrayFromDataBaseOwnerToInsuranceCompany:=<-channel_DOIC_insuranceCompany_SendEncryptedFile

								time.Sleep(time.Second*5)

								err=json.Unmarshal(byteArrayFromDataBaseOwnerToInsuranceCompany, &DataBaseOwnerToInsuranceCompany_object)
								if err != nil {
											log.Fatal(err)
								}

								MR_encryptedFileChunks:=createMerkleRootForInputVectors(DataBaseOwnerToInsuranceCompany_object.EncryptedOutputOfGates)


								dbDetails_object,err:=instance_P_IC_DBO.ClaimIDToDBODetails(nil,big.NewInt(int64(claimId)))
								if err != nil {
											log.Fatal(err)
								}


								bool_verifyingSignatureOfDb_SigEncFile:=VerifySignatue(dbDetails_object.DboAddress,MR_encryptedFileChunks,dbDetails_object.SignatureDBOHashOfEncryptedFile)


			insuranceCompany_GettingEncFileAndVerifyingDbSignature_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time InsuranceCompany_GettingEncFileAndVerifyingDbSignature : %v \n\n", insuranceCompany_GettingEncFileAndVerifyingDbSignature_stop.Sub(insuranceCompany_GettingEncFileAndVerifyingDbSignature_start))


			if bool_verifyingSignatureOfDb_SigEncFile{



										insuranceCompany_lockingApprovedMoney_start:=time.Now()


															claimDetails_object,err:=instance_P_IC_DBO.ClaimDetails(nil,big.NewInt(int64(claimId-1)))
															if err!=nil{
																	log.Fatal(err)
															}

															patientIC_claimedAmount:=claimDetails_object.ClaimedAmount



															tx,err=instance_P_IC_DBO.LockClaimedMoney(&bind.TransactOpts{
																	From:insuranceCompany.From,
																	Signer:insuranceCompany.Signer,
																	Value:patientIC_claimedAmount,
															},con_address_registration,big.NewInt(int64(claimId)),big.NewInt(int64(patientId)))
															if err!=nil{
																	log.Fatal(err)
															}
															client.Commit()

										insuranceCompany_lockingApprovedMoney_stop:=time.Now()
										log.SetPrefix(prefix)
										log.Printf("Time InsuranceCompany_lockingClaimeddMoney : %v \n\n", insuranceCompany_lockingApprovedMoney_stop.Sub(insuranceCompany_lockingApprovedMoney_start))
										log.SetPrefix(prefix)
										log.Println("Tx InsuranceCompany_LockClaimedMoney : ",tx.Gas(),"\n")




										insuranceCompany_waitingForPatientToRevealSecretKey_start:=time.Now()


															ic_ApproveClaimMoney:=false


															for i:=0;i<timeLimitForKeyReveal;i++ {


																		claimDetails_object,err:=instance_P_IC_DBO.ClaimDetails(nil,big.NewInt(int64(claimId-1)))
																		if err!=nil{
																				log.Fatal(err)
																		}


																		TimestampKeyReveal:=claimDetails_object.TimestampKeyReveal
																		if err!=nil{
																				log.Fatal(err)
																		}

																		if TimestampKeyReveal.Int64()!=0	{
																			ic_ApproveClaimMoney=true
																			break
																		}

																		time.Sleep(time.Second *1)

														 }


										insuranceCompany_waitingForPatientToRevealSecretKey_stop:=time.Now()
										log.SetPrefix(prefix)
										log.Printf("Time InsuranceCompany_waitingForPatientToRevealSecretKey : %v \n\n", insuranceCompany_waitingForPatientToRevealSecretKey_stop.Sub(insuranceCompany_waitingForPatientToRevealSecretKey_start))


										if !ic_ApproveClaimMoney{

															insuranceCompany_UnlockingAmount_patientNotRevealedKey_start:=time.Now()


																				claimDetails_object,err:=instance_P_IC_DBO.ClaimDetails(nil,big.NewInt(int64(claimId-1)))
																				if err!=nil{
																						log.Fatal(err)
																				}



																				TimestampIcLocking:=claimDetails_object.TimestampIcLocking

																				Now,err:=instance_P_IC_DBO.Now(nil)
																				if err!=nil{
																							log.Fatal(err)
																				}



																				for {


																									if Now.Int64()>TimestampIcLocking.Int64()+ timeLimitForKeyReveal{
																										break
																									}

																									_,err:=instance_P_IC_DBO.IncreaseTheTimeBy10seconds(&bind.TransactOpts{
																												From:insuranceCompany.From,
																												Signer:insuranceCompany.Signer,
																									})
																									if err!=nil{
																											log.Fatal(err)
																									}
																									client.Commit()

																									Now,err=instance_P_IC_DBO.Now(nil)
																									if err!=nil{
																											log.Fatal(err)
																									}



																				}


																				tx,err=instance_P_IC_DBO.WithdrawingLockedClaimedMoney(&bind.TransactOpts{
																							From:insuranceCompany.From,
																							Signer:insuranceCompany.Signer,
																				},con_address_registration,big.NewInt(int64(claimId)),big.NewInt(int64(patientId)))
																				if err!=nil{
																						log.Fatal(err)
																				}
																				client.Commit()



															insuranceCompany_UnlockingAmount_patientNotRevealedKey_stop:=time.Now()
															log.SetPrefix(prefix)
															log.Printf("Time InsuranceCompany_UnlockingAmount_patientNotRevealedKey : %v \n\n", insuranceCompany_UnlockingAmount_patientNotRevealedKey_stop.Sub(insuranceCompany_UnlockingAmount_patientNotRevealedKey_start))
															log.SetPrefix(prefix)
															log.Println("Tx InsuranceCompany_LeaveAfterLockedMoney : ",tx.Gas(),"\n")


															wg.Done()

										}


										claimDetails_object,err=instance_P_IC_DBO.ClaimDetails(nil,big.NewInt(int64(claimId-1)))
										if err!=nil{
												log.Fatal(err)
										}

										secret_key:=claimDetails_object.SecretKey
										rsa_secretKey:=BytesToPrivateKey(secret_key)
										fileChunks:=DecryptFileChunks(DataBaseOwnerToInsuranceCompany_object.EncryptedOutputOfGates,rsa_secretKey)
										MR_filechunks:=createMerkleRootForInputVectors(fileChunks)
										estimatedBillId:=claimDetails_object.EstimatedBillID


										multiSigID,err:=instance_P_HA_1.EstimatedCostBillIDToMultiSigOnMedicalDataID(nil,estimatedBillId)
										if err!=nil{
												log.Fatal(err)
										}

										multiSigID_int:=multiSigID.Int64()


										MultiSigOnMedicalData_object,err:=instance_P_HA_1.MultiSigOnMedicalData(nil,big.NewInt(int64(multiSigID_int-1)))
										if err!=nil{
												log.Fatal(err)
										}



										if MultiSigOnMedicalData_object.HashOfData!=MR_filechunks{

														insuranceCompany_ApproveClaimedMoney_start:=time.Now()


																		approvedAmount:=0

																		tx,err=instance_P_IC_DBO.ApproveClaim(&bind.TransactOpts{
																					From:insuranceCompany.From,
																					Signer:insuranceCompany.Signer,
																					Value:big.NewInt(int64(approvedAmount)),
																		},con_address_registration,big.NewInt(int64(claimId)),big.NewInt(int64(patientId)),big.NewInt(int64(approvedAmount)))
																		if err!=nil{
																				log.Fatal(err)
																		}
																		client.Commit()


														insuranceCompany_ApproveClaimedMoney_stop:=time.Now()
														log.SetPrefix(prefix)
														log.Printf("Time InsuranceCompany_ApproveClaimedMoney_zero : %v \n\n", insuranceCompany_ApproveClaimedMoney_stop.Sub(insuranceCompany_ApproveClaimedMoney_start))
														log.SetPrefix(prefix)
														log.Println("Tx InsuranceCompany_ApproveClaimedMoney_zero : ",tx.Gas(),"\n")

														wg.Done()


										}



										insuranceCompany_ApproveClaimedMoney_start:=time.Now()


														approvedAmount:=30

														tx,err=instance_P_IC_DBO.ApproveClaim(&bind.TransactOpts{
																	From:insuranceCompany.From,
																	Signer:insuranceCompany.Signer,
																	Value:big.NewInt(int64(approvedAmount)),
														},con_address_registration,big.NewInt(int64(claimId)),big.NewInt(int64(patientId)),big.NewInt(int64(approvedAmount)))
														if err!=nil{
																log.Fatal(err)
														}
														client.Commit()


										insuranceCompany_ApproveClaimedMoney_stop:=time.Now()
										log.SetPrefix(prefix)
										log.Printf("Time InsuranceCompany_ApproveClaimedMoney : %v \n\n", insuranceCompany_ApproveClaimedMoney_stop.Sub(insuranceCompany_ApproveClaimedMoney_start))
										log.SetPrefix(prefix)
										log.Println("Tx InsuranceCompany_ApproveClaimedMoney : ",tx.Gas(),"\n")



										wg.Done()

			} else {

					fmt.Println("\n\nIC - DB's signature doesnt match")
					wg.Done()

			}






	}


	func EncryptFileChunks(_fileChunks [][]byte, publicKey *rsa.PublicKey) ([][]byte){

			encryptedFileChunks:=make([][]byte,noOfInputGates)

			var i uint

				for i=0;i<noOfInputGates;i++{
					encryptedFileChunks[i]=rsa_PublicKey_Encrypt(_fileChunks[i],publicKey)
				}
			return encryptedFileChunks

	}

	func DecryptFileChunks(_encryptedFileChunks [][]byte, privateKey *rsa.PrivateKey) ([][]byte){

				decryptedFileChunks:=make([][]byte,noOfInputGates)

				var i uint

				for i=0;i<noOfInputGates;i++{
					decryptedFileChunks[i]=rsa_PrivateKey_Decrypt(_encryptedFileChunks[i],privateKey)
				}
				return decryptedFileChunks


	}



	func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {


				privkey, err := rsa.GenerateKey(rand.Reader, bits)
				if err != nil {
					log.Fatal(err)
				}
				return privkey, &privkey.PublicKey
  }

		// PrivateKeyToBytes private key to bytes
		func PrivateKeyToBytes(priv *rsa.PrivateKey) []byte {

					privBytes := pem.EncodeToMemory(
						&pem.Block{
							Type:  "RSA PRIVATE KEY",
							Bytes: x509.MarshalPKCS1PrivateKey(priv),
						},
					)

					return privBytes
		}

			// PublicKeyToBytes public key to bytes
			func PublicKeyToBytes(pub *rsa.PublicKey) []byte {

					pubASN1, err := x509.MarshalPKIXPublicKey(pub)
					if err != nil {
						log.Fatal(err)
					}

					pubBytes := pem.EncodeToMemory(&pem.Block{
						Type:  "RSA PUBLIC KEY",
						Bytes: pubASN1,
					})

					return pubBytes
			}

			// BytesToPrivateKey bytes to private key
			func BytesToPrivateKey(priv []byte) *rsa.PrivateKey {

					block, _ := pem.Decode(priv)
					enc := x509.IsEncryptedPEMBlock(block)
					b := block.Bytes
					var err error

					if enc {
						log.Println("is encrypted pem block")
						b, err = x509.DecryptPEMBlock(block, nil)
						if err != nil {
							log.Fatal(err)
						}
					}

					key, err := x509.ParsePKCS1PrivateKey(b)
					if err != nil {
						log.Fatal(err)
					}
					return key

			}

			// BytesToPublicKey bytes to public key
			func BytesToPublicKey(pub []byte) *rsa.PublicKey {

						block, _ := pem.Decode(pub)
						enc := x509.IsEncryptedPEMBlock(block)
						b := block.Bytes
						var err error


						if enc {
							log.Println("is encrypted pem block")
							b, err = x509.DecryptPEMBlock(block, nil)
							if err != nil {
								log.Fatal(err)
							}
						}


						ifc, err := x509.ParsePKIXPublicKey(b)
						if err != nil {
							log.Fatal(err)
						}

						key, ok := ifc.(*rsa.PublicKey)
						if !ok {
							log.Fatal("not ok")
						}
						return key
			}

			func rsa_PublicKey_Encrypt(plainText []byte,publicKey *rsa.PublicKey) ([]byte) {

		        label := []byte("")

		        encryptedtext, err := rsa.EncryptOAEP(
		            sha256.New(),
		            rand.Reader,
		            publicKey,
		            plainText,
		            label,
		        )
		        if err!=nil{
		         	log.Println(err)
		        }

		        return encryptedtext

			}

			func rsa_PrivateKey_Decrypt(encryptedtext []byte,privateKey *rsa.PrivateKey) ([]byte){

						label := []byte("")

			      plainText, err := rsa.DecryptOAEP(
			    		sha256.New(),
			        rand.Reader,
			        privateKey,
			        encryptedtext,
			        label,
			    	)
			      if err!=nil{
			        log.Println(err)
			      }
						return plainText

			}









































		func Patient_Do(patient_DO *bind.TransactOpts,client *backends.SimulatedBackend,channel_Patient_DO_appId chan int,channel_Patient_DO_EncryptedFile_Object chan []byte,channel_Patient_DO_DBOid chan int,wg *sync.WaitGroup){

				prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : Patient_WithDataBaseOwner "

				instance_P_DO,err:=import_contract_p_dbo.NewSCPDBO(con_address_P_DO,client)
				instance_registration,err								:=import_contract_registration.NewSystemUsersInfo(con_address_registration,client)


				var tx *types.Transaction
				var PatientToDataBaseOwner_Object PatientToDataBaseOwner

				patientId:=1
				estimatedBillId:=1

				dboId_int:=<-channel_Patient_DO_DBOid



				dboAddress,_,_,err:=instance_registration.DboById(nil,big.NewInt(int64(dboId_int)))
				if err!=nil{
						log.Fatal(err)
				}




				patientDO_keepData_start:=time.Now()

								key:=keyGenerate()
								inputVectors:=FileRead()
								MRinp:=createMerkleRootForInputVectors(inputVectors)
								encodedGateOutputs,_:=Encode(inputVectors,key,MRinp)
								MRencout:=createMerkleRootForEncInp(encodedGateOutputs)

								tx,err=instance_P_DO.KeepDataPatient(&bind.TransactOpts{
										From:patient_DO.From,
										Signer:patient_DO.Signer,
										Value:big.NewInt(20),
								},con_address_registration,con_address_P_HA_1,big.NewInt(int64(estimatedBillId)),MRencout,MRinp,dboAddress)
								if err!=nil{
								    log.Fatal(err)
								}
								client.Commit()


				patientDO_keepData_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientDO_keepData : %v \n\n", patientDO_keepData_stop.Sub(patientDO_keepData_start))
				log.SetPrefix(prefix)
				log.Println("Tx KeepDataPatient : ",tx.Gas(),"\n")



				patientDO_sendingAppIdToDO_start:=time.Now()

								app_Id,err:=instance_P_DO.PatientIdToApplicationId(nil,big.NewInt(int64(patientId)))
								if err!=nil{
										log.Fatal(err)
								}

								PatientToDataBaseOwner_Object.EncryptedOutputOfGates=encodedGateOutputs
								byteArrayFromPatientToDateBaseOwner,err := json.Marshal(PatientToDataBaseOwner_Object)
								if err != nil {
											log.Fatal(err)
								}


								channel_Patient_DO_appId<-int(app_Id.Int64())
								channel_Patient_DO_EncryptedFile_Object<-byteArrayFromPatientToDateBaseOwner



				patientDO_sendingAppIdToDO_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientDO_sendingAppIdToDO : %v \n\n", patientDO_sendingAppIdToDO_stop.Sub(patientDO_sendingAppIdToDO_start))


				patientDO_WaitingForDbToVerifyEncFile_start:=time.Now()

								patientDO_DBVerifyingEncFile:=false

								for i:=0;i<timeLimitForMerkleRootVerification_P_DO;i++ {

									ApplicationForStoring_Object,err:=instance_P_DO.ApplicationForStoring(nil,big.NewInt(int64(app_Id.Int64()-1)))
									if err!=nil{
											log.Fatal(err)
									}

									if ApplicationForStoring_Object.TimestampMerkleRootVerification.Int64()!=0	{
										patientDO_DBVerifyingEncFile=true
										break
									}
									time.Sleep(time.Second *1)

								}


				patientDO_WaitingForDbToVerifyEncFile_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientDO_WaitingForDbToVerifyEncFile : %v \n\n", patientDO_WaitingForDbToVerifyEncFile_stop.Sub(patientDO_WaitingForDbToVerifyEncFile_start))


				if !patientDO_DBVerifyingEncFile{

						  patientDO_UnlockingAmount_DBNotVerifiedEncFile_start:=time.Now()

											ApplicationForStoring_Object,err:=instance_P_DO.ApplicationForStoring(nil,big.NewInt(int64(app_Id.Int64()-1)))
											if err!=nil{
													log.Fatal(err)
											}


											TimestampApplication:=ApplicationForStoring_Object.TimestampApplication

											Now,err:=instance_P_DO.Now(nil)


											for  {


																	if Now.Int64()>TimestampApplication.Int64()+timeLimitForMerkleRootVerification_P_DO {
																		break
																	}

																	_,err:=instance_P_DO.IncreaseTheTimeBy10seconds(&bind.TransactOpts{
																				From:patient_DO.From,
																				Signer:patient_DO.Signer,
																	})
																	if err!=nil{
																			log.Fatal(err)
																	}
																	client.Commit()

																	Now,err=instance_P_DO.Now(nil)
																	if err!=nil{
																			log.Fatal(err)
																	}


											}


											tx,err=instance_P_DO.LeaveAfterKeepDataPatient(&bind.TransactOpts{
													From:patient_DO.From,
													Signer:patient_DO.Signer,
											},con_address_registration,con_address_P_HA_1,app_Id,big.NewInt(int64(estimatedBillId)))
											if err!=nil{
											    log.Fatal(err)
											}
											client.Commit()


						patientDO_UnlockingAmount_DBNotVerifiedEncFile_stop:=time.Now()
						log.SetPrefix(prefix)
						log.Printf("Time PatientDO_UnlockingAmount_DBNotVerifiedEncFile : %v \n\n", patientDO_UnlockingAmount_DBNotVerifiedEncFile_stop.Sub(patientDO_UnlockingAmount_DBNotVerifiedEncFile_start))
						log.SetPrefix(prefix)
						log.Println("Tx PatientDO_LeaveAfterKeepDataPatient : ",tx.Gas(),"\n")

						wg.Done()

				}


				patientDO_KeyReveal_start:=time.Now()

									tx,err=instance_P_DO.KeyRevealPatient(&bind.TransactOpts{
											From:patient_DO.From,
											Signer:patient_DO.Signer,
									},con_address_registration,con_address_P_HA_1,app_Id,key)
									if err!=nil{
											log.Fatal(err)
									}
									client.Commit()


				patientDO_KeyReveal_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientDO_KeyReveal : %v \n\n", patientDO_KeyReveal_stop.Sub(patientDO_KeyReveal_start))
				log.SetPrefix(prefix)
				log.Println("Tx PatientDO_KeyReveal : ",tx.Gas(),"\n")



				patientDO_WaitingForDbToGiveApprovalOrComplain_start:=time.Now()

								patientDO_ApprovalOrComplain:=false

								for i:=0;i<timeLimitForGivingComplainOrPuttingInTheDatatbase_P_DO;i++ {

									ApplicationForStoring_Object,err:=instance_P_DO.ApplicationForStoring(nil,big.NewInt(int64(app_Id.Int64()-1)))
									if err!=nil{
											log.Fatal(err)
									}

									if ApplicationForStoring_Object.TimestampComplain.Int64()!=0	{
										patientDO_ApprovalOrComplain=true
										break
									}

									if ApplicationForStoring_Object.TimestampApproval.Int64()!=0	{
										patientDO_ApprovalOrComplain=true
										break
									}

									time.Sleep(time.Second *1)

								}


				patientDO_WaitingForDbToGiveApprovalOrComplain_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientDO_WaitingForDbToGiveApprovalOrComplain : %v \n\n", patientDO_WaitingForDbToGiveApprovalOrComplain_stop.Sub(patientDO_WaitingForDbToGiveApprovalOrComplain_start))


				if !patientDO_ApprovalOrComplain{

						  patientDO_UnlockingAmount_DBNotGivenApprovalOrComplain_start:=time.Now()

											ApplicationForStoring_Object,err:=instance_P_DO.ApplicationForStoring(nil,big.NewInt(int64(app_Id.Int64()-1)))
											if err!=nil{
													log.Fatal(err)
											}


											TimestampKeyReveal:=ApplicationForStoring_Object.TimestampKeyReveal

											Now,err:=instance_P_DO.Now(nil)


											for  {


																	if Now.Int64()>TimestampKeyReveal.Int64()+timeLimitForGivingComplainOrPuttingInTheDatatbase_P_DO {
																		break
																	}

																	_,err:=instance_P_DO.IncreaseTheTimeBy10seconds(&bind.TransactOpts{
																				From:patient_DO.From,
																				Signer:patient_DO.Signer,
																	})
																	if err!=nil{
																			log.Fatal(err)
																	}
																	client.Commit()

																	Now,err=instance_P_DO.Now(nil)
																	if err!=nil{
																			log.Fatal(err)
																	}


											}


											tx,err=instance_P_DO.LeaveAfterKeyRevealPatient(&bind.TransactOpts{
													From:patient_DO.From,
													Signer:patient_DO.Signer,
											},con_address_registration,con_address_P_HA_1,app_Id,big.NewInt(int64(estimatedBillId)))
											if err!=nil{
											    log.Fatal(err)
											}
											client.Commit()


						patientDO_UnlockingAmount_DBNotGivenApprovalOrComplain_stop:=time.Now()
						log.SetPrefix(prefix)
						log.Printf("Time PatientDO_WaitingForDbToGiveApprovalOrComplain : %v \n\n", patientDO_UnlockingAmount_DBNotGivenApprovalOrComplain_stop.Sub(patientDO_UnlockingAmount_DBNotGivenApprovalOrComplain_start))
						log.SetPrefix(prefix)
						log.Println("Tx PatientDO_LeaveAfterKeyRevealPatient : ",tx.Gas(),"\n")
						wg.Done()

				}


				wg.Done()

			}

		func DatabaseOwner(dataBaseOwner *bind.TransactOpts,client *backends.SimulatedBackend,channel_Patient_DO_appId chan int,channel_Patient_DO_EncryptedFile_Object chan []byte,channel_Patient_DO_DBOid chan int,wg *sync.WaitGroup){

				prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : DataBaseOwner "

				instance_P_DO,err:=import_contract_p_dbo.NewSCPDBO(con_address_P_DO,client)
				instance_registration,err								:=import_contract_registration.NewSystemUsersInfo(con_address_registration,client)


				var tx *types.Transaction
				var PatientToDataBaseOwner_Object PatientToDataBaseOwner


				_,dboId,_,err:=instance_registration.DboByAddress(nil,dataBaseOwner.From)
				if err!=nil{
						log.Fatal(err)
				}


				channel_Patient_DO_DBOid<-int(dboId.Int64())


				dataBaseOwner_waitingForAppIdAndEncFile_start:=time.Now()

							app_Id:=<-channel_Patient_DO_appId
							byteArrayFromPatientToDateBaseOwner:=<-channel_Patient_DO_EncryptedFile_Object

							err=json.Unmarshal(byteArrayFromPatientToDateBaseOwner, &PatientToDataBaseOwner_Object)
							if err != nil {
									log.Fatal(err)
							}


				dataBaseOwner_waitingForAppIdAndEncFile_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time DataBaseOwner_waitingForAppIdAndEncFile : %v \n\n", dataBaseOwner_waitingForAppIdAndEncFile_stop.Sub(dataBaseOwner_waitingForAppIdAndEncFile_start))

				ApplicationForStoring_Object,err:=instance_P_DO.ApplicationForStoring(nil,big.NewInt(int64(app_Id-1)))
				if err != nil {
						log.Fatal(err)
				}


				if !checkMerkleRootForEncInp(PatientToDataBaseOwner_Object.EncryptedOutputOfGates,ApplicationForStoring_Object.MerkleRootOfEncFile){
					wg.Done()
				}


				dataBaseOwner_verifyingEncFile_start:=time.Now()

									tx,err=instance_P_DO.VerifyEncFileDO(&bind.TransactOpts{
											From:dataBaseOwner.From,
											Signer:dataBaseOwner.Signer,
											Value:big.NewInt(20),
									},big.NewInt(int64(app_Id)))
									if err!=nil{
											log.Fatal(err)
									}
									client.Commit()


				dataBaseOwner_verifyingEncFile_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time DataBaseOwner_VerifyingEncFile : %v \n\n", dataBaseOwner_verifyingEncFile_stop.Sub(dataBaseOwner_verifyingEncFile_start))
				log.SetPrefix(prefix)
				log.Println("Tx DataBaseOwnerVerifyingEncFile : ",tx.Gas(),"\n")



				dataBaseOwner_waitingForPatientToKeyReveal_start:=time.Now()

								dataBaseOwner_keyReveal:=false

								for i:=0;i<timeLimitForKeyReveal_P_DO;i++ {

									ApplicationForStoring_Object,err:=instance_P_DO.ApplicationForStoring(nil,big.NewInt(int64(app_Id-1)))
									if err!=nil{
											log.Fatal(err)
									}

									if ApplicationForStoring_Object.TimestampKeyReveal.Int64()!=0	{
										dataBaseOwner_keyReveal=true
										break
									}
									time.Sleep(time.Second *1)

								}


				dataBaseOwner_waitingForPatientToKeyReveal_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time DataBaseOwner_waitingForPatientToKeyReveal : %v \n\n", dataBaseOwner_waitingForPatientToKeyReveal_stop.Sub(dataBaseOwner_waitingForPatientToKeyReveal_start))



				if !dataBaseOwner_keyReveal{

								dataBaseOwner_UnlockingPatientNotRevealedKey_start:=time.Now()

												ApplicationForStoring_Object,err:=instance_P_DO.ApplicationForStoring(nil,big.NewInt(int64(app_Id-1)))
												if err!=nil{
														log.Fatal(err)
												}


												TimestampMerkleRootVerification:=ApplicationForStoring_Object.TimestampMerkleRootVerification

												Now,err:=instance_P_DO.Now(nil)


												for  {


																	if Now.Int64()>TimestampMerkleRootVerification.Int64()+timeLimitForKeyReveal_P_DO {
																		break
																	}

																	_,err:=instance_P_DO.IncreaseTheTimeBy10seconds(&bind.TransactOpts{
																				From:dataBaseOwner.From,
																				Signer:dataBaseOwner.Signer,
																	})
																	if err!=nil{
																			log.Fatal(err)
																	}
																	client.Commit()

																	Now,err=instance_P_DO.Now(nil)
																	if err!=nil{
																			log.Fatal(err)
																	}


											}


											tx,err=instance_P_DO.LeaveAfterVerifyEncFileDO(&bind.TransactOpts{
													From:dataBaseOwner.From,
													Signer:dataBaseOwner.Signer,
											},big.NewInt(int64(app_Id)))
											if err!=nil{
													log.Fatal(err)
											}
											client.Commit()



							dataBaseOwner_UnlockingPatientNotRevealedKey_stop:=time.Now()
							log.SetPrefix(prefix)
							log.Printf("Time DataBaseOwner_UnlockingPatientNotRevealedKey : %v \n\n", dataBaseOwner_UnlockingPatientNotRevealedKey_stop.Sub(dataBaseOwner_UnlockingPatientNotRevealedKey_start))
							log.SetPrefix(prefix)
							log.Println("Tx DataBaseOwnerLeaveAfterVerifyEncFileDO : ",tx.Gas(),"\n")


							wg.Done()


				}


				dataBaseOwner_decodingEncodedVectors_start:=time.Now()


										ApplicationForStoring_Object,err=instance_P_DO.ApplicationForStoring(nil,big.NewInt(int64(app_Id-1)))
										if err!=nil{
												log.Fatal(err)
										}

										key:=ApplicationForStoring_Object.Key

										dataBaseOwner_merkleTreeForEncInput:=createMerkleTreeForEncInp(PatientToDataBaseOwner_Object.EncryptedOutputOfGates)

										complain,decodedOutputs,encodedVectors_ThisGate,gateIndexes:=Extract(PatientToDataBaseOwner_Object.EncryptedOutputOfGates,key,dataBaseOwner_merkleTreeForEncInput,ApplicationForStoring_Object.MerkleRootOfEncFile,ApplicationForStoring_Object.MerkleRootOfFile)



				dataBaseOwner_decodingEncodedVectors_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time DataBaseOwner_decodingEncodedVectors : %v \n\n", dataBaseOwner_decodingEncodedVectors_stop.Sub(dataBaseOwner_decodingEncodedVectors_start))


				if complain==nil{


						fmt.Println("decodedOutputs : ",decodedOutputs,"\n")

						dataBaseOwner_approval_start:=time.Now()

									tx,err:=instance_P_DO.ApprovalDO(&bind.TransactOpts{
											From:dataBaseOwner.From,
											Signer:dataBaseOwner.Signer,
									},big.NewInt(int64(app_Id)))
									if err!=nil{
											log.Fatal(err)
									}
									client.Commit()


						dataBaseOwner_approval_stop:=time.Now()
						log.SetPrefix(prefix)
						log.Printf("Time DataBaseOwner_approval : %v \n\n", dataBaseOwner_approval_stop.Sub(dataBaseOwner_approval_start))
						log.SetPrefix(prefix)
						log.Println("Tx DataBaseOwner_approval : ",tx.Gas(),"\n")

				}



				if complain!=nil{



								dataBaseOwner_Complain_start:=time.Now()



											tx,err:=instance_P_DO.Complain(&bind.TransactOpts{
													From:dataBaseOwner.From,
													Signer:dataBaseOwner.Signer,
											},con_address_registration,con_address_P_HA_1,complain,encodedVectors_ThisGate,big.NewInt(int64(gateIndexes[0])),big.NewInt(int64(app_Id)))
											if err!=nil{
													log.Fatal(err)
											}
											client.Commit()




							dataBaseOwner_Complain_stop:=time.Now()
							log.SetPrefix(prefix)
							log.Printf("Time DataBaseOwner_Complain : %v \n\n", dataBaseOwner_Complain_stop.Sub(dataBaseOwner_Complain_start))
							log.SetPrefix(prefix)
							log.Println("Tx DataBaseOwner_Complain : ",tx.Gas(),"\n")

							wg.Done()

						}


						wg.Done()


	}



	func Patient(patient *bind.TransactOpts,client *backends.SimulatedBackend,channel_HospitalToPatient_EstimatedBillId chan int,channel_HospitalToPatient_EncryptedFile_Object chan []byte,channel_PatientToHospital_PatientId chan int,channel_HospitalToPatient_HospitalId chan int,channel_multiSignId chan int,wg *sync.WaitGroup){

		prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : Patient "

		//contract instances
		instance_registration,err:=import_contract_p_dbo.NewSystemUsersInfo(con_address_registration,client)
		if err!=nil{
		    log.Fatal(err)
		}

		instance_P_HA_1,err:=import_contract_p_dbo.NewSCPHA1(con_address_P_HA_1,client)
		if err!=nil{
		    log.Fatal(err)
		}

		instance_P_HA_2,err:=import_contract_p_dbo.NewSCPHA2(con_address_P_HA_2,client)
		if err!=nil{
		    log.Fatal(err)
		}

		var tx *types.Transaction

		//patient registration
		_,err=instance_registration.PatientRegistration(&bind.TransactOpts{
				From:patient.From,
				Signer:patient.Signer,
		},"amd",big.NewInt(20),big.NewInt(8123456789),"ISI")
		if err!=nil{
				log.Fatal(err)
		}
		client.Commit()

		_,patientId,_,_,_,_,err:=instance_registration.GetPatientbyAddress(nil,patient.From)
		if err!=nil{
				log.Fatal(err)
		}

		fmt.Println("\npatientId : ",patientId,"\n")


		hospitalId:=<-channel_HospitalToPatient_HospitalId
		channel_PatientToHospital_PatientId<-int(patientId.Int64())
		estimatedBillId:=<-channel_HospitalToPatient_EstimatedBillId






			/*Patient Locking*/
			patient_LockEstimatedAmount_start := time.Now()

								tx,err=instance_P_HA_1.LockEstimatedAmount(&bind.TransactOpts{
										From:patient.From,
										Signer:patient.Signer,
									  Value:big.NewInt(50),
								},con_address_registration,big.NewInt(int64(hospitalId)),big.NewInt(int64(estimatedBillId)))
								if err!=nil{
										log.Fatal(err)
								}
								client.Commit()

			patient_LockEstimatedAmount_stop := time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time Patient_LockingEstimatedAmount : %v \n\n", patient_LockEstimatedAmount_stop.Sub(patient_LockEstimatedAmount_start))

			log.SetPrefix(prefix)
			log.Println("Tx Patient_LockingEstimatedAmount : ",tx.Gas(),"\n")




			/*Waiting For Hospital To start Treatment*/
			patient_WaitingForHospitalToStartTreatment_start:=time.Now()


							patient_hospitalStartingTreatment:=false

							for i:=0;i<timeLimitForStartingOperation;i++ {

								estimatedCheckUpCost_Object,err:=instance_P_HA_1.EstimatedCheckUpCost(nil,big.NewInt(int64(estimatedBillId-1)))
								if err!=nil{
										log.Fatal(err)
								}

								if estimatedCheckUpCost_Object.TimestampOfCheckUpStart.Int64()!=0	{
									patient_hospitalStartingTreatment=true
									break
								}
								time.Sleep(time.Second *1)

							}

			patient_WaitingForHospitalToStartTreatment_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time Patient_WaitingForHospitalToStartTreatment : %v \n\n", patient_WaitingForHospitalToStartTreatment_stop.Sub(patient_WaitingForHospitalToStartTreatment_start))



			if !patient_hospitalStartingTreatment{

					patient_UnlockingEstimatedAmount_HospitalNotStartedTreatment_start:=time.Now()



									estimatedCheckUpCost_Object,err:=instance_P_HA_1.EstimatedCheckUpCost(nil,big.NewInt(int64(estimatedBillId-1)))
									if err!=nil{
											log.Fatal(err)
									}

									TimestampOfLockingByPatient:=estimatedCheckUpCost_Object.TimestampOfLockingByPatient

									Now,err:=instance_P_HA_1.Now(nil)

									for  {


												if Now.Int64()>TimestampOfLockingByPatient.Int64()+timeLimitForStartingOperation {
													break
												}

												_,err:=instance_P_HA_1.IncreaseTheTimeBy10seconds(&bind.TransactOpts{
															From:patient.From,
															Signer:patient.Signer,
												})
												if err!=nil{
														log.Fatal(err)
												}
												client.Commit()

												Now,err=instance_P_HA_1.Now(nil)
												if err!=nil{
														log.Fatal(err)
												}


									}

									tx,err=instance_P_HA_2.PatientLeaveLockEstimatedAmount(&bind.TransactOpts{
												From:patient.From,
												Signer:patient.Signer,
										},con_address_registration,con_address_P_HA_1,big.NewInt(int64(hospitalId)),big.NewInt(int64(estimatedBillId)))
										if err!=nil{
												log.Fatal(err)
										}

										client.Commit()

						patient_UnlockingEstimatedAmount_HospitalNotStartedTreatment_stop:=time.Now()
						log.SetPrefix(prefix)
						log.Printf("Time Patient_UnlockingEstimatedAmount_HospitalNotStartedTreatment : %v \n\n", patient_UnlockingEstimatedAmount_HospitalNotStartedTreatment_stop.Sub(patient_UnlockingEstimatedAmount_HospitalNotStartedTreatment_start))
						log.SetPrefix(prefix)
						log.Println("Tx PatientLeaveLockEstimatedAmount : ",tx.Gas(),"\n")


						wg.Done()



			}


				patient_WaitingForHospitalToSendMultiSignIdAndEncryptedFiles_start:=time.Now()


							multiSigID:=<-channel_multiSignId


							patient_MultiSigOnMedicalDataObject,err:=instance_P_HA_1.MultiSigOnMedicalData(nil,big.NewInt(int64(multiSigID-1)))
							if err!=nil{
									log.Fatal(err)
							}


							var patient_HospitalToPatientObject HospitalToPatient
							byteArrayFromHospitalToPatient:=<-channel_HospitalToPatient_EncryptedFile_Object



				patient_WaitingForHospitalToSendMultiSignIdAndEncryptedFiles_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time Patient_WaitingForHospitalToSendMultiSignIdAndEncryptedFiles : %v \n\n", patient_WaitingForHospitalToSendMultiSignIdAndEncryptedFiles_stop.Sub(patient_WaitingForHospitalToSendMultiSignIdAndEncryptedFiles_start))






				patient_CheckingMerkleRootAndVerifySignature_start:=time.Now()


							err=json.Unmarshal(byteArrayFromHospitalToPatient, &patient_HospitalToPatientObject)
							if err != nil {
									log.Fatal(err)
							}

							patient_hashDateMerkleRoot:=patient_MultiSigOnMedicalDataObject.HashOfPatientIdDateHashEncFile
							patient_signature_hashDateMerkleRoot:=patient_MultiSigOnMedicalDataObject.SignOnHashOfPatientIdDateOfReportHashOfEncryptedFile
							verification_hashDatePatient_signature:=VerifySignatue(patient_HospitalToPatientObject.HospitalAddress,patient_hashDateMerkleRoot,patient_signature_hashDateMerkleRoot)


							patient_contract_hashOfEncFile:=patient_MultiSigOnMedicalDataObject.HashOfEncryptedData
							bool_verification_merkleRoots_encFile:=checkMerkleRootForEncInp(patient_HospitalToPatientObject.EncryptedOutputOfGates,patient_contract_hashOfEncFile)



			 patient_CheckingMerkleRootAndVerifySignature_stop:=time.Now()
			 log.SetPrefix(prefix)
			 log.Printf("Time Patient_CheckingMerkleRootAndVerifySignature : %v \n\n", patient_CheckingMerkleRootAndVerifySignature_stop.Sub(patient_CheckingMerkleRootAndVerifySignature_start))





			if bool_verification_merkleRoots_encFile && verification_hashDatePatient_signature{

						patient_VerifyingSignatureAndMerkleRoot_start:=time.Now()


									tx,err=instance_P_HA_1.VerifyAndGiveConsent(&bind.TransactOpts{
												From:patient.From,
												Signer:patient.Signer,
									},con_address_registration,big.NewInt(int64(multiSigID)),big.NewInt(int64(hospitalId)),patient_hashDateMerkleRoot,patient_MultiSigOnMedicalDataObject.HashOfEncryptedData)
									if err!=nil{
											log.Fatal(err)
									}
									client.Commit()




						patient_VerifyingSignatureAndMerkleRoot_stop:=time.Now()
						log.SetPrefix(prefix)
		 			  log.Printf("Time Patient_VerifyingSignatureAndMerkleRoot : %v \n\n", patient_VerifyingSignatureAndMerkleRoot_stop.Sub(patient_VerifyingSignatureAndMerkleRoot_start))
						log.SetPrefix(prefix)
						log.Println("Tx PatientVerifyAndGiveConsent : ",tx.Gas(),"\n")

			}




			patient_WaitingForHospitalToCommitFinalBill_start:=time.Now()



								var finalBillID *big.Int
								patient_hospitalGeneratingFinalBill:=false

								for i:=0;i<timeLimitForGeneratingFillBill;i++ {

										finalBillID,err=instance_P_HA_1.EstimatedBillIDToFinalBillID(nil,big.NewInt(int64(estimatedBillId)))
										if err!=nil{
												log.Fatal(err)
										}

										if finalBillID.Int64()!=0	{
											patient_hospitalGeneratingFinalBill=true
											break
										}
										time.Sleep(time.Second *1)

								}



			patient_WaitingForHospitalToCommitFinalBill_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time Patient_WaitingForHospitalToCommitFinalBill : %v \n\n", patient_WaitingForHospitalToCommitFinalBill_stop.Sub(patient_WaitingForHospitalToCommitFinalBill_start))


			if !patient_hospitalGeneratingFinalBill{

					  patient_LeavingSinceHospitalNotCommitedFinalBill_start:=time.Now()


										MultiSigOnMedicalData_Object,err:=instance_P_HA_1.MultiSigOnMedicalData(nil,big.NewInt(int64(multiSigID-1)))
										if err != nil {
											log.Fatal(err)
										}


										TimeStampOfVerificationByPatient:=MultiSigOnMedicalData_Object.TimeStampOfVerificationByPatient

										Now,err:=instance_P_HA_1.Now(nil)

										for  {


													if Now.Int64()>TimeStampOfVerificationByPatient.Int64()+timeLimitForGeneratingFillBill {
														break
													}

													_,err:=instance_P_HA_1.IncreaseTheTimeBy10seconds(&bind.TransactOpts{
																From:patient.From,
																Signer:patient.Signer,
													})
													if err!=nil{
															log.Fatal(err)
													}
													client.Commit()

													Now,err=instance_P_HA_1.Now(nil)
													if err!=nil{
															log.Fatal(err)
													}


										}

										tx,err=instance_P_HA_2.PatientLeaveConsentFinalBill(&bind.TransactOpts{
													From:patient.From,
													Signer:patient.Signer,
											},con_address_registration,con_address_P_HA_1,big.NewInt(int64(estimatedBillId)),big.NewInt(int64(hospitalId)))
											if err!=nil{
													log.Fatal(err)
											}

										client.Commit()





						patient_LeavingSinceHospitalNotCommitedFinalBill_stop:=time.Now()
						log.SetPrefix(prefix)
						log.Printf("Time Patient_LeavingSinceHospitalNotCommitedFinalBill : %v \n\n", patient_LeavingSinceHospitalNotCommitedFinalBill_stop.Sub(patient_LeavingSinceHospitalNotCommitedFinalBill_start))
						log.SetPrefix(prefix)
						log.Println("Tx PatientLeaveConsentFinalBill : ",tx.Gas(),"\n")

						wg.Done()

			}


			patient_GiveConsentFinalBill_start:=time.Now()

							tx,err=instance_P_HA_1.ConsentFinalBillPatient(&bind.TransactOpts{
										From:patient.From,
										Signer:patient.Signer,
							},con_address_registration,finalBillID,big.NewInt(int64(hospitalId)))
							if err!=nil{
									log.Fatal(err)
							}
							client.Commit()


			patient_GiveConsentFinalBill_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time Patient_GiveConsentFinalBill : %v \n\n", patient_GiveConsentFinalBill_stop.Sub(patient_GiveConsentFinalBill_start))
			log.SetPrefix(prefix)
			log.Println("Tx PatientGiveConsentFinalBill : ",tx.Gas(),"\n")



			patient_WaitingForHospitalToRevealKey_start:=time.Now()

							patient_hospitalKeyReveal:=false

							for i:=0;i<timeLimitForKeyRevealing;i++ {

									keyAndExchange_object,err:=instance_P_HA_1.MultiSignIdToKeyAndExchange(nil,big.NewInt(int64(multiSigID)))
									if err!=nil{
											log.Fatal(err)
									}

									if (keyAndExchange_object.TimeStampOfKeyReveal.Int64())!=0	{
										patient_hospitalKeyReveal=true
										break
									}
									time.Sleep(time.Second *1)

							}


			patient_WaitingForHospitalToRevealKey_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time Patient_WaitingForHospitalToRevealKey : %v \n\n", patient_WaitingForHospitalToRevealKey_stop.Sub(patient_WaitingForHospitalToRevealKey_start))



			if !patient_hospitalKeyReveal{

					patient_LeavingSinceHospitalNotRevealedKey_start:=time.Now()


									finalBill_Object,err:=instance_P_HA_1.FinalCheckUpCost(nil,big.NewInt(finalBillID.Int64()-1))
									if err != nil {
										log.Fatal(err)
									}


									TimeStampOfPatientFinalBillConsent:=finalBill_Object.TimeStampOfPatientFinalBillConsent

									Now,err:=instance_P_HA_1.Now(nil)

									for  {


												if Now.Int64()>TimeStampOfPatientFinalBillConsent.Int64()+timeLimitForKeyRevealing {
													break
												}

												_,err:=instance_P_HA_1.IncreaseTheTimeBy10seconds(&bind.TransactOpts{
															From:patient.From,
															Signer:patient.Signer,
												})
												if err!=nil{
														log.Fatal(err)
												}
												client.Commit()

												Now,err=instance_P_HA_1.Now(nil)
												if err!=nil{
														log.Fatal(err)
												}


									}

									tx,err=instance_P_HA_2.PatientLeaveConsentFinalBill(&bind.TransactOpts{
											From:patient.From,
											Signer:patient.Signer,
									},con_address_registration,con_address_P_HA_1,big.NewInt(int64(estimatedBillId)),big.NewInt(int64(hospitalId)))
									if err!=nil{
											log.Fatal(err)
									}

									client.Commit()

									wg.Done()



					patient_LeavingSinceHospitalNotRevealedKey_stop:=time.Now()
					log.SetPrefix(prefix)
					log.Printf("Time Patient_LeavingSinceHospitalNotRevealedKey : %v \n\n", patient_LeavingSinceHospitalNotRevealedKey_stop.Sub(patient_LeavingSinceHospitalNotRevealedKey_start))
					log.SetPrefix(prefix)
					log.Println("Tx PatientLeaveAfterGivingConsentFinalBill : ",tx.Gas(),"\n")



			}


			patient_DecryptingEncodedGateOutputs_start:=time.Now()


							keyAndExchange_object,err:=instance_P_HA_1.MultiSignIdToKeyAndExchange(nil,big.NewInt(int64(multiSigID)))
							if err!=nil{
									log.Fatal(err)
							}


							key:=keyAndExchange_object.Key

							patient_merkleTreeForEncInput:=createMerkleTreeForEncInp(patient_HospitalToPatientObject.EncryptedOutputOfGates)

							complain,decodedOutputs,encodedVectors_ThisGate,gateIndexes:=Extract(patient_HospitalToPatientObject.EncryptedOutputOfGates,key,patient_merkleTreeForEncInput,patient_contract_hashOfEncFile,patient_hashDateMerkleRoot)



			patient_DecryptingEncodedGateOutputs_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time Patient_DecryptingEncodedGateOutputs : %v \n\n", patient_DecryptingEncodedGateOutputs_stop.Sub(patient_DecryptingEncodedGateOutputs_start))



			if complain==nil{


					fmt.Println("decodedOutputs : ",decodedOutputs,"\n")

					patient_FinalConsent_start:=time.Now()

								tx,err:=instance_P_HA_1.SatisfyAndCompleteTheProcess(&bind.TransactOpts{
										From:patient.From,
										Signer:patient.Signer,
								},con_address_registration,big.NewInt(int64(estimatedBillId)),big.NewInt(int64(hospitalId)))
								if err!=nil{
										log.Fatal(err)
								}
								client.Commit()


					patient_FinalConsent_stop:=time.Now()
					log.SetPrefix(prefix)
					log.Printf("Time Patient_FinalConsent : %v \n\n", patient_FinalConsent_stop.Sub(patient_FinalConsent_start))
					log.SetPrefix(prefix)
					log.Println("Tx PatientFinalConsent : ",tx.Gas(),"\n")




					// fmt.Println("\n\nGave consent")
			}

				if complain !=nil{



								patient_Complain_start:=time.Now()


											var gateIndexes_bigInt []*big.Int

											for i:=0;i<len(gateIndexes);i++{
												gateIndexes_bigInt=append(gateIndexes_bigInt,big.NewInt(int64(gateIndexes[i])))

											}


											tx,err:=instance_P_HA_1.PatientComplain(&bind.TransactOpts{
													From:patient.From,
													Signer:patient.Signer,
											},con_address_registration,complain,encodedVectors_ThisGate,gateIndexes_bigInt,big.NewInt(int64(estimatedBillId)),big.NewInt(int64(hospitalId)))
											if err!=nil{
													log.Fatal(err)
											}
											client.Commit()


							patient_Complain_stop:=time.Now()
							log.SetPrefix(prefix)
							log.Printf("Time Patient_Complain : %v \n\n", patient_Complain_stop.Sub(patient_Complain_start))
							log.SetPrefix(prefix)
							log.Println("Tx PatientFinalConsent : ",tx.Gas(),"\n")

							wg.Done()

						}


			wg.Done()


	}

	func Hospital(hospital *bind.TransactOpts,client *backends.SimulatedBackend,channel_HospitalToPatient_EncryptedFile_Object chan []byte,channel_HospitalToPatient_EstimatedBillId chan int,channel_PatientToHospital_PatientId chan int,channel_HospitalToPatient_HospitalId chan int,channel_HospitalToPatient_sendHashDateMRenc chan [hashFunctionOutputBitSize]byte,channel_multiSignId chan int,wg *sync.WaitGroup,hospitalprivateKey *ecdsa.PrivateKey){

				prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : Hospital "

				estimateBillCost:=50
				finalBillCost:=50

				//contract instances
				instance_registration,err:=import_contract_p_dbo.NewSystemUsersInfo(con_address_registration,client)
				if err!=nil{
				    log.Fatal(err)
				}

				instance_P_HA_1,err:=import_contract_p_dbo.NewSCPHA1(con_address_P_HA_1,client)
				if err!=nil{
				    log.Fatal(err)
				}

				instance_P_HA_2,err:=import_contract_p_dbo.NewSCPHA2(con_address_P_HA_2,client)
				if err!=nil{
				    log.Fatal(err)
				}

				var tx *types.Transaction
				var HospitalToPatientObject HospitalToPatient



				//Hospital registration
				_,err=instance_registration.HospitalRegistration(&bind.TransactOpts{
						From:hospital.From,
						Signer:hospital.Signer,
				},"Apollo")
				if err!=nil{
						log.Fatal(err)
				}
				client.Commit()


				_,hospitalId,_,err:=instance_registration.GetHospitalbyAddress(nil,hospital.From)
				if err!=nil{
						log.Fatal(err)
				}

				fmt.Println("\nhospitalId : ",hospitalId,"\n")

				channel_HospitalToPatient_HospitalId<-int(hospitalId.Int64())
				patientId:=<-channel_PatientToHospital_PatientId




				//Hospital generates estimate bill
				hospital_GenerateEstimateBill_start := time.Now()

						  tx,err=instance_P_HA_1.GenerateEstimatedCostBill(&bind.TransactOpts{
									From:hospital.From,
									Signer:hospital.Signer,
									Value:big.NewInt(1000),
							},con_address_registration,big.NewInt(int64(patientId)),big.NewInt(int64(estimateBillCost)))
							if err!=nil{
									log.Fatal(err)
							}
							client.Commit()

				hospital_GenerateEstimateBill_stop := time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time hospital_GenerateEstimateBill : %v \n\n", hospital_GenerateEstimateBill_stop.Sub(hospital_GenerateEstimateBill_start))
				log.SetPrefix(prefix)
				log.Println("Tx HospitalGenerateEstimateBill : ",tx.Gas(),"\n")





				hospital_SendingEstimateBill_start := time.Now()

							estimatedBillId,err:=instance_P_HA_1.PatientIdHospitalIdEstimateBillId(nil,big.NewInt(int64(patientId)),hospitalId)
							if err!=nil{
									log.Fatal(err)
							}

							estimatedBillId_int:=estimatedBillId.Int64()


							channel_HospitalToPatient_EstimatedBillId<-int(estimatedBillId.Int64())

				hospital_SendingEstimateBill_stop := time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time Hospital_SendingEstimateBill : %v \n\n", hospital_SendingEstimateBill_stop.Sub(hospital_SendingEstimateBill_start))






				hospital_WaitingForPatientToLockEstimateAmount_start:=time.Now()

								hospital_lockingByPatient:=false


								for i:=0;i<timeLimitForPatientLocking;i++{

										estimatedCheckUpCost_Object,err:=instance_P_HA_1.EstimatedCheckUpCost(nil,big.NewInt(estimatedBillId_int-1))
										if err!=nil{
												log.Fatal(err)
										}


										if (estimatedCheckUpCost_Object.TimestampOfLockingByPatient).Int64()!=0	{
											hospital_lockingByPatient=true
											break
										}
										time.Sleep(time.Second *1)
								}


				hospital_WaitingForPatientToLockEstimateAmount_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time Hospital_WaitingForPatientToLockEstimateAmount : %v \n\n", hospital_WaitingForPatientToLockEstimateAmount_stop.Sub(hospital_WaitingForPatientToLockEstimateAmount_start))


				if !hospital_lockingByPatient {

						hospital_LeavingSincePatientNotLockedEstimateAmount_start:=time.Now()



											estimatedCheckUpCost_Object,err:=instance_P_HA_1.EstimatedCheckUpCost(nil,big.NewInt(estimatedBillId_int-1))
											if err!=nil{
													log.Fatal(err)
											}

											TimestampOfLockingByHospital:=estimatedCheckUpCost_Object.TimestampOfLockingByHospital

											Now,err:=instance_P_HA_1.Now(nil)

											for  {


														if Now.Int64()>TimestampOfLockingByHospital.Int64()+timeLimitForPatientLocking {
															break
														}

														_,err:=instance_P_HA_1.IncreaseTheTimeBy10seconds(&bind.TransactOpts{
																	From:hospital.From,
																	Signer:hospital.Signer,
														})
														if err!=nil{
																log.Fatal(err)
														}
														client.Commit()

														Now,err=instance_P_HA_1.Now(nil)
														if err!=nil{
																log.Fatal(err)
														}

											}





											tx,err=instance_P_HA_2.HospitalLeaveGenerateEstimatedCostBill(&bind.TransactOpts{
														From:hospital.From,
														Signer:hospital.Signer,
											},con_address_registration,con_address_P_HA_1,big.NewInt(int64(patientId)),estimatedBillId)
											if err!=nil{
													log.Fatal(err)
											}

											client.Commit()



						hospital_LeavingSincePatientNotLockedEstimateAmount_stop:=time.Now()
						log.SetPrefix(prefix)
						log.Printf("Time Hospital_LeavingSincePatientNotLockedEstimateAmount : %v \n\n", hospital_LeavingSincePatientNotLockedEstimateAmount_stop.Sub(hospital_LeavingSincePatientNotLockedEstimateAmount_start))
						log.SetPrefix(prefix)
						log.Println("Tx HospitalLeaveGenerateEstimatedCostBill : ",tx.Gas(),"\n")

						wg.Done()
				}



				hospital_StartTreatment_start:=time.Now()


							tx,err=instance_P_HA_1.StartTreatment(&bind.TransactOpts{
										From:hospital.From,
										Signer:hospital.Signer,
								},con_address_registration,big.NewInt(int64(patientId)),estimatedBillId)
								if err!=nil{
										log.Fatal(err)
								}
								client.Commit()


				hospital_StartTreatment_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time Hospital_StartTreatment : %v \n\n", hospital_StartTreatment_stop.Sub(hospital_StartTreatment_start))
				log.SetPrefix(prefix)
				log.Println("Tx HospitalStartTreatment : ",tx.Gas(),"\n")



				hospital_CalculatingVariables_start:=time.Now()


										key:=keyGenerate()
										keyCommit:=fnKeycommit(key)
										inputVectors:=FileRead()
										MRinp:=createMerkleRootForInputVectors(inputVectors)
										encodedGateOutputs,_:=Encode(inputVectors,key,MRinp)
										MRencout:=createMerkleRootForEncInp(encodedGateOutputs)



										bytefileHash:=make([]byte,32)
										for i:=0;i<32;i++{
												bytefileHash[i]=MRinp[i]
										}

										signature_hospital_p1_fileHash, err := crypto.Sign(bytefileHash,hospitalprivateKey)
										if err != nil {
											log.Fatal(err)
										}


										var total_date []int

										date:=29
										month:=3
										year:=2020
										total_date=append(total_date,date,month,year)


										var dateMerkleRoot []byte


										dateMerkleRoot=append(dateMerkleRoot,byte(patientId))

										for i:=0;i<len(total_date);i++ {

											dateMerkleRoot=append(dateMerkleRoot,byte(total_date[i]))

										}



										for i:=0;i<32;i++{
											dateMerkleRoot=append(dateMerkleRoot,MRencout[i])
										}
										hashDateMerkleRoot:=Hash(dateMerkleRoot)
										byteHashDateMerkleRoot:=make([]byte,32)
										for i:=0;i<32;i++{
												byteHashDateMerkleRoot[i]=hashDateMerkleRoot[i]
										}




										signature_hospital_p1_patientIdDate, err := crypto.Sign(byteHashDateMerkleRoot,hospitalprivateKey)
										if err != nil {
											log.Fatal(err)
										}



				hospital_CalculatingVariables_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time Hospital_CalculatingVariables : %v \n\n", hospital_CalculatingVariables_stop.Sub(hospital_CalculatingVariables_start))




		    hospital_KeepingSignedHashToBlockchain_start:=time.Now()


							_,err=instance_P_HA_1.KeepSignedHashToBlockchain(&bind.TransactOpts{
										From:hospital.From,
										Signer:hospital.Signer,
								},con_address_registration,big.NewInt(int64(patientId)),estimatedBillId,MRinp,MRencout,hashDateMerkleRoot,signature_hospital_p1_fileHash,signature_hospital_p1_patientIdDate,big.NewInt(int64(noOfInputGates)),big.NewInt(int64(buffer_size)),big.NewInt(maxLinesToGate),keyCommit)
								if err!=nil{
										log.Fatal(err)
								}
								client.Commit()


				hospital_KeepingSignedHashToBlockchain_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time Hospital_KeepingSignedHashToBlockchain : %v \n\n", hospital_KeepingSignedHashToBlockchain_stop.Sub(hospital_KeepingSignedHashToBlockchain_start))
				log.SetPrefix(prefix)
				log.Println("Tx HospitalKeepSignedHashToBlockchain : ",tx.Gas(),"\n")



				hospital_SendingMultiSignIdAndEncodedOuptuts_start:=time.Now()


									multiSignId,err:=instance_P_HA_1.EstimatedCostBillIDToMultiSigOnMedicalDataID(nil,estimatedBillId)
									if err!=nil{
											log.Fatal(err)
									}

									multiSignId_int:=multiSignId.Int64()


									channel_multiSignId<-int(multiSignId_int)



									HospitalToPatientObject.EncryptedOutputOfGates=encodedGateOutputs
									HospitalToPatientObject.HospitalAddress=hospital.From

									byteArrayFromHospitalToPatient,err := json.Marshal(HospitalToPatientObject)
									if err != nil {
										log.Fatal(err)
									}

									channel_HospitalToPatient_EncryptedFile_Object<-byteArrayFromHospitalToPatient



				hospital_SendingMultiSignIdAndEncodedOuptuts_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time Hospital_SendingMultiSignIdAndEncodedOuptuts : %v \n\n", hospital_SendingMultiSignIdAndEncodedOuptuts_stop.Sub(hospital_SendingMultiSignIdAndEncodedOuptuts_start))




				hospital_WaitingForPatientToVerifySignatureAndMerkleRoot_start:=time.Now()


								hospital_VerifyingSignature:=false

								for i:=0;i<timeLimitForVerifyingSignature;i++ {

									multiSignObject,err:=instance_P_HA_1.MultiSigOnMedicalData(nil,big.NewInt(multiSignId_int-1))
									if err != nil {
										log.Fatal(err)
									}
									if multiSignObject.TimeStampOfVerificationByPatient.Int64()!=0	{
										hospital_VerifyingSignature=true
										break
									}

									time.Sleep(time.Second *1)
								}




				hospital_WaitingForPatientToVerifySignatureAndMerkleRoot_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time Hospital_WaitingForPatientToVerifySignatureAndMerkleRoot : %v \n\n", hospital_WaitingForPatientToVerifySignatureAndMerkleRoot_stop.Sub(hospital_WaitingForPatientToVerifySignatureAndMerkleRoot_start))


				// fmt.Println("\n\n hospital_lockingByPatient : ",hospital_lockingByPatient)

				if !hospital_VerifyingSignature {


							hospital_UnlockingEstimatedAmountSincePatientNotVerifySignature_start:=time.Now()


											MultiSigOnMedicalData_Object,err:=instance_P_HA_1.MultiSigOnMedicalData(nil,big.NewInt(multiSignId_int-1))
											if err != nil {
												log.Fatal(err)
											}

											TimeStampOfSignByHospital:=MultiSigOnMedicalData_Object.TimeStampOfSignByHospital

											Now,err:=instance_P_HA_1.Now(nil)
											if err != nil {
												log.Fatal(err)
											}

											for  {


														if Now.Int64()>TimeStampOfSignByHospital.Int64()+timeLimitForStartingOperation {
															break
														}

														_,err:=instance_P_HA_1.IncreaseTheTimeBy10seconds(&bind.TransactOpts{

																		From:hospital.From,
																		Signer:hospital.Signer,
														})
														if err!=nil{
																log.Fatal(err)
														}
														client.Commit()


														Now,err=instance_P_HA_1.Now(nil)
														if err!=nil{
																log.Fatal(err)
														}


										 }



											tx,err=instance_P_HA_2.HospitalLeaveKeepHashToBlockchain(&bind.TransactOpts{
														From:hospital.From,
														Signer:hospital.Signer,
											},con_address_registration,con_address_P_HA_1,big.NewInt(int64(patientId)),estimatedBillId)

											client.Commit()



							hospital_UnlockingEstimatedAmountSincePatientNotVerifySignature_stop:=time.Now()
							log.SetPrefix(prefix)
							log.Printf("Time Hospital_UnlockingEstimatedAmountSincePatientNotVerifySignature : %v \n\n", hospital_UnlockingEstimatedAmountSincePatientNotVerifySignature_stop.Sub(hospital_UnlockingEstimatedAmountSincePatientNotVerifySignature_start))
							log.SetPrefix(prefix)
							log.Println("Tx HospitalLeaveAfterKeepingHashToBlockchain : ",tx.Gas(),"\n")


							wg.Done()

				}

					hospital_DischargeAndGenerateFinalBill_start:=time.Now()


									tx,err=instance_P_HA_1.DischargeAndGenerateFinalCostBill(&bind.TransactOpts{
											From:hospital.From,
											Signer:hospital.Signer,
											// Value:big.NewInt(50),
									},con_address_registration,estimatedBillId,big.NewInt(int64(patientId)),big.NewInt(int64(finalBillCost)))
									if err!=nil{
											log.Fatal(err)
									}
									client.Commit()





					hospital_DischargeAndGenerateFinalBill_stop:=time.Now()
					log.SetPrefix(prefix)
					log.Printf("Time Hospital_DischargeAndGenerateFinalBill : %v \n\n", hospital_DischargeAndGenerateFinalBill_stop.Sub(hospital_DischargeAndGenerateFinalBill_start))
					log.SetPrefix(prefix)
					log.Println("Tx HospitalDischargeAndGenerateFinalCostBill : ",tx.Gas(),"\n")



					finalBillID,err:=instance_P_HA_1.EstimatedBillIDToFinalBillID(nil,estimatedBillId)
					if err!=nil{
							log.Fatal(err)
					}

					finalBillID_int:=finalBillID.Int64()


					hospital_WaitingForPatientToGiveConsentFinalBill_start:=time.Now()


									hospital_ConsentFinalBill:=false

									for i:=0;i<timeLimitForSigningFinalBill;i++ {

											finalBill_Object,err:=instance_P_HA_1.FinalCheckUpCost(nil,big.NewInt(finalBillID_int-1))
											if err != nil {
												log.Fatal(err)
											}
											if finalBill_Object.TimeStampOfPatientFinalBillConsent.Int64()!=0	{
												hospital_ConsentFinalBill=true
												break
											}

											time.Sleep(time.Second *1)

									}


					hospital_WaitingForPatientToGiveConsentFinalBill_stop:=time.Now()
					log.SetPrefix(prefix)
					log.Printf("Time Hospital_WaitingForPatientToGiveConsentFinalBill : %v \n\n", hospital_WaitingForPatientToGiveConsentFinalBill_stop.Sub(hospital_WaitingForPatientToGiveConsentFinalBill_start))



					if !hospital_ConsentFinalBill {


						hospital_UnlockingFinalAmountSincePatientNotConsent_start:=time.Now()


										finalBill_Object,err:=instance_P_HA_1.FinalCheckUpCost(nil,big.NewInt(finalBillID_int-1))
										if err != nil {
											log.Fatal(err)
										}


										TimestampOfFinalBilling:=finalBill_Object.TimestampOfFinalBilling

										Now,err:=instance_P_HA_1.Now(nil)
										if err != nil {
											log.Fatal(err)
										}

										for  {


												if Now.Int64()>TimestampOfFinalBilling.Int64()+timeLimitForSigningFinalBill {
													break
												}

												_,err:=instance_P_HA_1.IncreaseTheTimeBy10seconds(&bind.TransactOpts{

																From:hospital.From,
																Signer:hospital.Signer,
												})
												if err!=nil{
														log.Fatal(err)
												}
												client.Commit()


												Now,err=instance_P_HA_1.Now(nil)
												if err!=nil{
														log.Fatal(err)
												}


									}


									tx,err=instance_P_HA_2.HospitalLeaveDischargeAndGenerateFinalCostBill(&bind.TransactOpts{
												From:hospital.From,
												Signer:hospital.Signer,
									},con_address_registration,con_address_P_HA_1,estimatedBillId,big.NewInt(int64(patientId)))

									client.Commit()



						hospital_UnlockingFinalAmountSincePatientNotConsent_stop:=time.Now()
						log.SetPrefix(prefix)
						log.Printf("Time Hospital_UnlockingFinalAmountSincePatientNotConsent : %v \n\n", hospital_UnlockingFinalAmountSincePatientNotConsent_stop.Sub(hospital_UnlockingFinalAmountSincePatientNotConsent_start))
						log.SetPrefix(prefix)
						log.Println("Tx HospitalLeaveDischargeAndGenerateFinalCostBill : ",tx.Gas(),"\n")

						wg.Done()



				}


				hospital_KeyReveal_start:=time.Now()


							tx,err=instance_P_HA_1.RevealKey(&bind.TransactOpts{
										From:hospital.From,
										Signer:hospital.Signer,
							},con_address_registration,key,estimatedBillId,big.NewInt(int64(patientId)))
							if err!=nil{
									log.Fatal(err)
							}
							client.Commit()


			hospital_KeyReveal_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time Hospital_KeyReveal : %v \n\n", hospital_KeyReveal_stop.Sub(hospital_KeyReveal_start))
			log.SetPrefix(prefix)
			log.Println("Tx hospitalKeyReveal : ",tx.Gas(),"\n")



			wg.Done()


	}




	func VerifySignatue(address common.Address,hashOnWhichSig [hashFunctionOutputBitSize]byte,signature []byte) (bool){

						byte_hashOnWhichSig:=make([]byte,32)

						for i:=0;i<32;i++{
							byte_hashOnWhichSig[i]=hashOnWhichSig[i]
						}

						party_PublicKey, err := crypto.Ecrecover(byte_hashOnWhichSig, signature)
						if err != nil {
								log.Fatal(err)
						}

						party_PublicKeyECDSA,err:=crypto.UnmarshalPubkey(party_PublicKey)
						if err != nil {
								log.Fatal(err)
						}


					  party_Address:= crypto.PubkeyToAddress(*party_PublicKeyECDSA)

						if party_Address==address{
							return true
						}else {
							return false
						}


	}


  func Decrypt(key [keySize]byte,encryptedtext []byte) ([]byte){

				finalKey:=generateFinalKey(key,len(encryptedtext))
				var plainText []byte
				for i:=0;i<len(encryptedtext);i++{
					plainText=append(plainText,finalKey[i]^encryptedtext[i])
				}

				return plainText
	}

	func Encrypt(key [keySize]byte,plainText []byte) ([]byte) {

				finalKey:=generateFinalKey(key,len(plainText))
				var encryptedtext []byte
				for i:=0;i<len(plainText);i++{
					encryptedtext=append(encryptedtext,finalKey[i]^plainText[i])
				}
				return encryptedtext
	}



	func AuthAndAddressGeneration() (*bind.TransactOpts,common.Address,*ecdsa.PrivateKey){

		//privateKey *ecdsa.PrivateKey
	  privateKey,err:=crypto.GenerateKey()
	  if err!=nil{
	    log.Fatal(err)
	  }

	  auth:=bind.NewKeyedTransactor(privateKey)

	  publicKey:=privateKey.Public()
	  publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	  if !ok {
	      log.Fatal("error casting public key to ECDSA")
	  }
	  address := crypto.PubkeyToAddress(*publicKeyECDSA)
	  return auth,address,privateKey
	}



	func keyGenerate() ([keySize]byte){

		keyGen := make([]byte, keySize)
		var key [keySize]byte

	  _, err := rand.Read(keyGen)
	  if err != nil {
	      log.Fatal(err)
	  }

		for i:=0;i<len(keyGen);i++{
			key[i]=keyGen[i]
		}
		return key
	}



	func generateFinalKey(key [32]byte,keySizeNeeded int) ([]byte) {
		var keyHashed []byte
		var temp [32]byte
		var finalKey []byte
		finalKey = make([]byte, keySizeNeeded)

		for i:=0;i<keySizeNeeded/32;i++ {

			if(i==0){
				var keyPlusIndex []byte

				k:=0
				for j:=0;j<32;j++{
					keyPlusIndex=append(keyPlusIndex,key[j])
				}
				keyPlusIndex=append(keyPlusIndex,byte(k))
				keyHashed = kec.Keccak256(keyPlusIndex)

				for j:=0;j<32;j++{
					temp[j]=keyHashed[j]
					finalKey[j]=temp[j]
				}

			}else{

				keyHashed=nil

				for j:=0;j<32;j++{
					keyHashed=append(keyHashed,temp[j])
				}
				keyHashed=kec.Keccak256(keyHashed)

				for j:=0;j<32;j++{
					temp[j]=keyHashed[j]
					finalKey[i*32+j]=temp[j]
				}
			}
		}
		return finalKey
	}




	func Hash(_toBeHashed []byte) ([32]byte){

		var toBeHashed []byte
		var Hashed [32]byte

		for j:=0;j<len(_toBeHashed);j++{
			toBeHashed=append(toBeHashed,_toBeHashed[j])
		}

		HashedAfter:=kec.Keccak256(toBeHashed)

		for j:=0;j<32;j++{
				Hashed[j]=HashedAfter[j]
		}

		return Hashed

	}



	func FileRead() ([][]byte){
		var currentByte int64 = 0
		fileBuffer := make([]byte, buffer_size)
		FileChunks := make([][]byte, noOfInputGates)
		file, err := os.Open("input.txt")
		if err != nil {
				log.Fatal(err)
		}
		var count uint
		count=0
		for {
				n,err := file.ReadAt(fileBuffer, currentByte)
				if err!=nil{
					log.Println(err)
				}
				currentByte += int64(buffer_size)
				fileBuffer=fileBuffer[:n]
				if count==noOfInputGates{
						break
				}
				FileChunks[count] = make([]byte, buffer_size)
				copy(FileChunks[count],fileBuffer)
				if err == io.EOF {
						log.Fatal(err)
				}
				count++
		}
		file.Close()
		return FileChunks
	}



	func createMerkleRootForInputVectors(inputVectors [][]byte) ([hashFunctionOutputBitSize]byte){

		var merkletree [][]byte
		merkletree = make([][]byte,2*noOfInputGates-1)
		var i uint

		//input gates
		for i=0;i<noOfInputGates;i++{
					merkletree[i]=inputVectors[i]
		}

		//!inputgates
		k:=0

		for i:=noOfInputGates;i<noOfInputGates+noOfInputGates/2;i++{

			var toBeHashed []byte
			var Hashed []byte

			var j uint
			for j=0;j<buffer_size;j++{
				toBeHashed=append(toBeHashed,merkletree[k][j])
			}

			for j=0;j<buffer_size;j++{
				toBeHashed=append(toBeHashed,merkletree[k+1][j])
			}

			Hashed=kec.Keccak256(toBeHashed)

			for j:=0;j<hashFunctionOutputBitSize;j++{
				merkletree[i]=append(merkletree[i],Hashed[j])
			}

			k=k+2

		}


		for i:=noOfInputGates+noOfInputGates/2;i<2*noOfInputGates-1;i++{

			var toBeHashed []byte
			var Hashed []byte

			for j:=0;j<hashFunctionOutputBitSize;j++{
				toBeHashed=append(toBeHashed,merkletree[k][j])
			}

			for j:=0;j<hashFunctionOutputBitSize;j++{
				toBeHashed=append(toBeHashed,merkletree[k+1][j])
			}

			Hashed=kec.Keccak256(toBeHashed)

			for j:=0;j<hashFunctionOutputBitSize;j++{
				merkletree[i]=append(merkletree[i],Hashed[j])
			}
			k=k+2
		}


		var merkleRoot [hashFunctionOutputBitSize]byte
		for j:=0;j<hashFunctionOutputBitSize;j++{
			merkleRoot[j]=merkletree[2*noOfInputGates-2][j]
		}
		return merkleRoot
	}

	func createMerkleRootForEncInp(encryptedGateOutputs [][]byte) ([hashFunctionOutputBitSize]byte){

		// var merkletree [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte
		merkletree:=make([][hashFunctionOutputBitSize]byte,2*totalNumberOfGates-1)
		var i uint
		//input gates
		//no hashing of the input vectors to the fn.
		for i=0;i<totalNumberOfGates;i++{
					merkletree[i]=Hash(encryptedGateOutputs[i])
		}

		//!inputgates
		//merkle tree construction
		k:=0
		for i=0;i<totalNumberOfGates-1;i++{

			var toBeHashed []byte
			var Hashed []byte
			for j:=0;j<hashFunctionOutputBitSize;j++{
				toBeHashed=append(toBeHashed,merkletree[k][j])
			}

			for j:=0;j<hashFunctionOutputBitSize;j++{
				toBeHashed=append(toBeHashed,merkletree[k+1][j])
			}

			Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input
			for j:=0;j<hashFunctionOutputBitSize;j++{
				merkletree[i+totalNumberOfGates][j]=Hashed[j]
			}
			k=k+2
		}

		  //returns the merkleroot
		return merkletree[2*totalNumberOfGates-2]
	}

	func fnKeycommit(key [keySize]byte) ([hashFunctionOutputBitSize]byte){

				var toBEHashedKey []byte
				var HashedKey [hashFunctionOutputBitSize]byte

				for j:=0;j<keySize;j++{
					toBEHashedKey=append(toBEHashedKey,key[j])
				}

				HashedAfter:=kec.Keccak256(toBEHashedKey)

				for j:=0;j<hashFunctionOutputBitSize;j++{
						HashedKey[j]=HashedAfter[j]
				}

				return HashedKey
	}


	func rsa_fnKeycommit(key []byte) ([hashFunctionOutputBitSize]byte){

				var toBEHashedKey []byte
				var HashedKey [hashFunctionOutputBitSize]byte

				for j:=0;j<len(key);j++{
					toBEHashedKey=append(toBEHashedKey,key[j])
				}

				HashedAfter:=kec.Keccak256(toBEHashedKey)

				for j:=0;j<hashFunctionOutputBitSize;j++{
						HashedKey[j]=HashedAfter[j]
				}

				return HashedKey
	}

	func malicious_Encode(inputVectors [][]byte,key [32]byte,MRx [hashFunctionOutputBitSize]byte) ([][]byte,[][]byte)  {

		out:=make([][]byte,totalNumberOfGates)
		z:=make([][]byte,totalNumberOfGates)
		var i uint

		for i=0;i<noOfInputGates;i++ {
				out[i]=inputVectors[i]
				temp:=Encrypt(key,out[i])
				z[i]=temp
		}

		//Not Input Gates
		//For each gate in this range,the input vectors to a particular gate are processed according to the operation and the ouput and encryption of this output is stored
		l:=0
		for i:=noOfInputGates;i<noOfInputGates+noOfInputGates/2;i++{
				Op:=2  //Operation of gate by the index "i"

				if i==noOfInputGates{
					Op=1
				}

				//Operation 2: The gate takes in the output of two of its children, concatenates the two inputs and hashes it.Later the output is encrypted.
				if(Op==2) {
						leftsibling:=l
						l++
						rightsibling:=l
						l++
						var toBeHashed []byte
						var j uint
						for j=0;j<buffer_size;j++ {
								toBeHashed= append(toBeHashed,out[leftsibling][j])
						}
						for j=0;j<buffer_size;j++ {
								toBeHashed= append(toBeHashed,out[rightsibling][j])
						}
						hash := kec.Keccak256(toBeHashed)
						out[i]=hash
				}

				//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
				if(Op==3) {
					for j:=0;j<hashFunctionOutputBitSize;j++{
						if (j==0){
								if(out[i-1][j]==MRx[j]){
									out[i]=append(out[i],1)
								} else {
									out[i]=append(out[i],0)
											 }
								} else {
									if(out[i-1][j]==MRx[j]){
										out[i]=append(out[i],0)
									}	else{
									out[i]=append(out[i],1)
									}
						}
					}
				}

				if(Op==1) {
						var j uint
						for j=0;j<buffer_size;j++{
							out[i]=append(out[i],byte(0))
						}
				}

				temp:=Encrypt(key,out[i])
				z[i]=temp
			}


		for i:=noOfInputGates+noOfInputGates/2;i<totalNumberOfGates;i++{
				Op:=2  //Operation of gate by the index "i"
				if(i==totalNumberOfGates-1){
					Op=3
				}

				//Operation 2: The gate takes in the output of two of its children, concatenates the two inputs and hashes it.Later the output is encrypted.
				if(Op==2) {
						leftsibling:=l
						l++
						rightsibling:=l
						l++
						var toBeHashed []byte
						for j:=0;j<hashFunctionOutputBitSize;j++ {
								toBeHashed= append(toBeHashed,out[leftsibling][j])
						}
						for j:=0;j<hashFunctionOutputBitSize;j++ {
								toBeHashed= append(toBeHashed,out[rightsibling][j])
						}
						hash := kec.Keccak256(toBeHashed)
						for j:=0;j<hashFunctionOutputBitSize;j++{
							out[i]=append(out[i],hash[j])
						}
				}

				//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
				if(Op==3) {
					for j:=0;j<hashFunctionOutputBitSize;j++{
						if (j==0){
								if(out[i-1][j]==MRx[j]){
									out[i]=append(out[i],1)
								} else {
									out[i]=append(out[i],0)
											 }
								} else {
									if(out[i-1][j]==MRx[j]){
										out[i]=append(out[i],0)
									}	else{
									out[i]=append(out[i],1)
									}
						}
					}
				}

				if(Op==1) {
						for j:=0;j<hashFunctionOutputBitSize;j++{
							out[i]=append(out[i],byte(0))
						}
				}

				temp:=Encrypt(key,out[i])
				z[i]=temp
			}

		return z,out
	}


	func Encode(inputVectors [][]byte,key [32]byte,MRx [hashFunctionOutputBitSize]byte) ([][]byte,[][]byte)  {

		out:=make([][]byte,totalNumberOfGates)
		z:=make([][]byte,totalNumberOfGates)
		var i uint

		for i=0;i<noOfInputGates;i++ {
				out[i]=inputVectors[i]
				temp:=Encrypt(key,out[i])
				z[i]=temp
		}

		//Not Input Gates
		//For each gate in this range,the input vectors to a particular gate are processed according to the operation and the ouput and encryption of this output is stored
		l:=0
		for i:=noOfInputGates;i<noOfInputGates+noOfInputGates/2;i++{
				Op:=2  //Operation of gate by the index "i"

				//Operation 2: The gate takes in the output of two of its children, concatenates the two inputs and hashes it.Later the output is encrypted.
				if(Op==2) {
						leftsibling:=l
						l++
						rightsibling:=l
						l++
						var toBeHashed []byte
						var j uint
						for j=0;j<buffer_size;j++ {
								toBeHashed= append(toBeHashed,out[leftsibling][j])
						}
						for j=0;j<buffer_size;j++ {
								toBeHashed= append(toBeHashed,out[rightsibling][j])
						}
						hash := kec.Keccak256(toBeHashed)
						out[i]=hash
				}

				//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
				if(Op==3) {
					for j:=0;j<hashFunctionOutputBitSize;j++{
						if (j==0){
								if(out[i-1][j]==MRx[j]){
									out[i]=append(out[i],1)
								} else {
									out[i]=append(out[i],0)
											 }
								} else {
									if(out[i-1][j]==MRx[j]){
										out[i]=append(out[i],0)
									}	else{
									out[i]=append(out[i],1)
									}
						}
					}
				}

				if(Op==1) {
						var j uint
						for j=0;j<buffer_size;j++{
							out[i]=append(out[i],byte(0))
						}
				}

				temp:=Encrypt(key,out[i])
				z[i]=temp
			}


		for i:=noOfInputGates+noOfInputGates/2;i<totalNumberOfGates;i++{
				Op:=2  //Operation of gate by the index "i"
				if(i==totalNumberOfGates-1){
					Op=3
				}

				//Operation 2: The gate takes in the output of two of its children, concatenates the two inputs and hashes it.Later the output is encrypted.
				if(Op==2) {
						leftsibling:=l
						l++
						rightsibling:=l
						l++
						var toBeHashed []byte
						for j:=0;j<hashFunctionOutputBitSize;j++ {
								toBeHashed= append(toBeHashed,out[leftsibling][j])
						}
						for j:=0;j<hashFunctionOutputBitSize;j++ {
								toBeHashed= append(toBeHashed,out[rightsibling][j])
						}
						hash := kec.Keccak256(toBeHashed)
						for j:=0;j<hashFunctionOutputBitSize;j++{
							out[i]=append(out[i],hash[j])
						}
				}

				//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
				if(Op==3) {
					for j:=0;j<hashFunctionOutputBitSize;j++{
						if (j==0){
								if(out[i-1][j]==MRx[j]){
									out[i]=append(out[i],1)
								} else {
									out[i]=append(out[i],0)
											 }
								} else {
									if(out[i-1][j]==MRx[j]){
										out[i]=append(out[i],0)
									}	else{
									out[i]=append(out[i],1)
									}
						}
					}
				}

				if(Op==1) {
						for j:=0;j<hashFunctionOutputBitSize;j++{
							out[i]=append(out[i],byte(0))
						}
				}

				temp:=Encrypt(key,out[i])
				z[i]=temp
			}

		return z,out
	}


	func checkMerkleRootForEncInp(encryptedGateOutputs [][]byte,merklerootforEncInput [hashFunctionOutputBitSize]byte) (bool){

		// var merkletree [2*totalNumberOfGates-1][hashFunctionOutputBitSize]byte
		merkletree:=make([][hashFunctionOutputBitSize]byte,2*totalNumberOfGates-1)
		var i uint
		//input gates
		//no hashing of the input vectors to the fn.
		for i=0;i<totalNumberOfGates;i++{
					merkletree[i]=Hash(encryptedGateOutputs[i])
		}

		//!inputgates
		//merkle tree construction
		k:=0
		for i=0;i<totalNumberOfGates-1;i++{

			var toBeHashed []byte
			var Hashed []byte
			for j:=0;j<hashFunctionOutputBitSize;j++{
				toBeHashed=append(toBeHashed,merkletree[k][j])
			}

			for j:=0;j<hashFunctionOutputBitSize;j++{
				toBeHashed=append(toBeHashed,merkletree[k+1][j])
			}

			Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input
			for j:=0;j<hashFunctionOutputBitSize;j++{
				merkletree[i+totalNumberOfGates][j]=Hashed[j]
			}
			k=k+2
		}

		for j:=0;j<hashFunctionOutputBitSize;j++{
			if(merkletree[2*totalNumberOfGates-2][j]!=merklerootforEncInput[j]){
				return false
			}
		}

		return true

	}

	func Extract(encodedGateOutputs [][]byte,key [32]byte,merkleTreeForEncGateOutput [][hashFunctionOutputBitSize]byte,Receiver_MerkleRootOfEncInp [hashFunctionOutputBitSize]byte,Receiver_MerkleRootOfFileChunks [hashFunctionOutputBitSize]byte) ([][][hashFunctionOutputBitSize]byte,[][]byte,[][]byte,[]uint) {


			decodedOutputs:=make([][]byte,totalNumberOfGates)
			var encodedVectors_ThisGate  [][]byte
			 //to store all the decodedOutputs. Dynamic in size due to the fact that a wrong gate operation can lead to stopping the decoding process and
			//generating the complaint
			var complain [][][hashFunctionOutputBitSize]byte
			var out []byte
			var gateIndexes []uint

			//Decryption of input vectors. Have taken the consumption that an error cannot be generated in the input gates. Even if the sender has something fishy in the
			//input gates, this effect will be refelected in the next level of gates.So, a relevant complaint can be given to the contract.
			var i uint
			for i=0;i<noOfInputGates;i++ {
					out =Decrypt(key,encodedGateOutputs[i])
					for j:=0;j<len(out);j++{
						decodedOutputs[i] = append(decodedOutputs[i],out[j])
					}
			}

			//Decryption of not input gates - gates above the base level.
			var l uint
			l=0
			for i:=noOfInputGates;i<totalNumberOfGates;i++{

				out = Decrypt(key,encodedGateOutputs[i])

				for j:=0;j<len(out);j++{
					decodedOutputs[i] = append(decodedOutputs[i],out[j])
				}

				operation:=2
				if(i==totalNumberOfGates-1){
					operation=3
				}

				var operationInputs [][]byte
				operationInputs=append(operationInputs,decodedOutputs[l])
				operationInputs=append(operationInputs,decodedOutputs[l+1])

						 if(!reflect.DeepEqual(decodedOutputs[i],Operation(operation,operationInputs,decodedOutputs,Receiver_MerkleRootOfEncInp,Receiver_MerkleRootOfFileChunks)))	 {
						// OperationVal:=Operation(operation,operationInputs,decodedOutputs,Receiver_MerkleRootOfEncInp)
						// fmt.Println("OperationVal : ",OperationVal)

								// gateIndexes=append(gateIndexes,i)
								gateIndexes=append(gateIndexes,i,l,l+1)

								encodedVectors_ThisGate=append(encodedVectors_ThisGate,encodedGateOutputs[i])
								MproofTree:=Mproof(i,merkleTreeForEncGateOutput)
								complain=append(complain,MproofTree)

								//merkle proof for the inputs of a complaint gate and appending to the complain vector
								for m:=0;m<maxLinesToGate;m++{

										encodedVectors_ThisGate=append(encodedVectors_ThisGate,encodedGateOutputs[l+uint(m)])
										MproofTree:=Mproof(l+uint(m),merkleTreeForEncGateOutput)
										complain=append(complain,MproofTree)
								}
								return complain,decodedOutputs,encodedVectors_ThisGate,gateIndexes
							}
							l=l+2
					}

				return complain,decodedOutputs,encodedVectors_ThisGate,gateIndexes

	}


			func Mverify(complaint [][32]byte,_index int,ciphertextRoot [32]byte) (bool){

			_value:=complaint[0]
			depth:=uint(len(complaint))
			var Hashed []byte
			var i uint
			for i=0; i<depth-1; i++{
				if ((_index & (1<<i))>>i == 1)	{

					var toBeHashed []byte
					for j:=0;j<32;j++{
						toBeHashed=append(toBeHashed,complaint[i+1][j])
					}
					for j:=0;j<32;j++{
						toBeHashed=append(toBeHashed, _value[j])

					}

					Hashed = kec.Keccak256(toBeHashed)

					for j:=0;j<32;j++{
						_value[j]=Hashed[j]

					}

				}else{

					var toBeHashed []byte
					for j:=0;j<32;j++{
						toBeHashed=append(toBeHashed,_value[j])
					}
					for j:=0;j<32;j++{
						toBeHashed=append(toBeHashed,complaint[i+1][j])
					}

					Hashed = kec.Keccak256(toBeHashed)

					for j:=0;j<32;j++{
						_value[j]=Hashed[j]
					}

				}

			}
			return (_value==ciphertextRoot)
		}


	func Mproof(index uint,merkletree [][hashFunctionOutputBitSize]byte) ([][hashFunctionOutputBitSize]byte){

		depth:=int(mathlib.Log2(float64(totalNumberOfGates)))+1
		tree:=make([][32]byte,depth)
		tree[0]=merkletree[index]
		for i:=1;i<depth;i++{
			if(index%2==0){
				tree[i]=merkletree[index+1]
				index=index/2+totalNumberOfGates
			} else {
				tree[i]=merkletree[index-1]
				index=index/2+totalNumberOfGates
			}
		}
		return tree

	}

	func Operation(Op int,operationInputs [][]byte,decodedOutputs [][]byte,Receiver_MerkleRootOfEncInp [hashFunctionOutputBitSize]byte,Receiver_MerkleRootOfFileChunks [hashFunctionOutputBitSize]byte)  ([]byte){

			var result []byte
			if(Op==2) {
					leftsibling:=operationInputs[0]
					rightsibling:=operationInputs[1]
					var toBeHashed []byte
					for j:=0;j<len(leftsibling);j++ {
							toBeHashed= append(toBeHashed,leftsibling[j])
					}
					for j:=0;j<len(rightsibling);j++ {
							toBeHashed= append(toBeHashed,rightsibling[j])
					}

					hashed := kec.Keccak256(toBeHashed)

					for j:=0;j<hashFunctionOutputBitSize;j++{
						result=append(result,hashed[j])
					}
				}

				if(Op==4) {
						toBeHashed:=operationInputs[0]
						hashed := kec.Keccak256(toBeHashed)

						for j:=0;j<hashFunctionOutputBitSize;j++{
							if (j==0){
									if(hashed[j]==Receiver_MerkleRootOfFileChunks[j]){
										result=append(result,1)
									} else {
										result=append(result,0)
												 }
									} else {
										if(hashed[j]==Receiver_MerkleRootOfFileChunks[j]){
											result=append(result,0)
										}	else{
											result=append(result,1)
										}
							}
					}
				}

			//Operation 3: The gate checks the equivalency of the MRx(input parameter) and the input to the gate.
			if(Op==3) {
				merkleTreeOfDecryptedInputVectors:=make([][]byte,noOfInputGates)
				var k uint
				for k=0;k<noOfInputGates;k++{
							merkleTreeOfDecryptedInputVectors[k]=decodedOutputs[k]
				}
				merkleRootofDecodedInputVectors:=createMerkleRootForInputVectors(merkleTreeOfDecryptedInputVectors)
				for j:=0;j<hashFunctionOutputBitSize;j++{
					if (j==0){
							if(operationInputs[0][j]==merkleRootofDecodedInputVectors[j]){
								result=append(result,1)
							} else {
								result=append(result,0)
										 }
							} else {
								if(operationInputs[0][j]==merkleRootofDecodedInputVectors[j]){
									result=append(result,0)
								}	else{
									result=append(result,1)
								}
					}
				}
			}
			return result
	}

	func createMerkleTreeForEncInp(encryptedGateOutputs [][]byte) ([][hashFunctionOutputBitSize]byte){

		merkletree:=make([][hashFunctionOutputBitSize]byte,2*totalNumberOfGates-1)

		//input gates
		//no hashing of the input vectors to the fn.
		var i uint
		for i=0;i<totalNumberOfGates;i++{
					merkletree[i]=Hash(encryptedGateOutputs[i])
		}

		//!inputgates
		//merkle tree construction
		k:=0
		for i=0;i<totalNumberOfGates-1;i++{

			var toBeHashed []byte
			var Hashed []byte
			for j:=0;j<hashFunctionOutputBitSize;j++{
				toBeHashed=append(toBeHashed,merkletree[k][j])
			}

			for j:=0;j<hashFunctionOutputBitSize;j++{
				toBeHashed=append(toBeHashed,merkletree[k+1][j])
			}

			Hashed=kec.Keccak256(toBeHashed) //Hashing of concatenated input

			for j:=0;j<hashFunctionOutputBitSize;j++{
				merkletree[i+totalNumberOfGates][j]=Hashed[j]
			}
			k=k+2
		}

		  //returns the merkleroot
		return merkletree
	}
