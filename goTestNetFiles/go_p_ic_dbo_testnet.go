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
		"context"
		"crypto/rand"
		"crypto/rsa"
	  "crypto/ecdsa"
	 	"github.com/ethereum/go-ethereum/ethclient"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
		"github.com/ethereum/go-ethereum/crypto"
	  "github.com/ethereum/go-ethereum/common"
		"github.com/ethereum/go-ethereum/core/types"
		"crypto/sha256"
		"math/big"
		"crypto/x509"
    "encoding/pem"
		 mathlib "math"
		 import_contract_p_h "../compiledGoFiles_testNet/contract_p_h"
		 // import_contract_p_dbo "../compiledGoFiles_testNet/contract_p_dbo"
		 import_contract_p_ic_dbo "../compiledGoFiles_testNet/contract_p_ic_dbo"

)


var buffer_size uint
var noOfInputGates uint
var totalNumberOfGates uint


var con_address_registration common.Address
var con_address_P_HA common.Address
var con_address_P_DBO common.Address
var con_address_P_IC common.Address


const keySize=32
const maxLinesToGate =2
const hashFunctionOutputBitSize=32


const timeLimiitForPolicyBuying=100
const timeLimitForICLocking=100
const timeLimitForKeyReveal=100
const timeLimitForFinalApproval=100


type PatientICToDataBaseOwner struct {
		InsuranceComapanyId int
		EstimatedBillId int
		ClaimId int
}

type DataBaseOwnerToInsuranceCompany struct {
		EncryptedOutputOfGates [][]byte
}


func main(){


			noOfInputGates=32
			buffer_size=32
			totalNumberOfGates=2*noOfInputGates

			logFilePath:="../log_testNet/log_p_ic_dbo_testNet.txt"

			var err error
			f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
					// log.Fatalf("error opening file: %v", err)
			defer f.Close()
			}
			wrt := io.MultiWriter(os.Stdout, f)
			log.SetOutput(wrt)




			ctx := context.Background()

			client, err := ethclient.Dial("Infura Key Point")
			if err != nil {
				log.Fatal(err)
			}


			patient_PrivateKey,err:=crypto.HexToECDSA("Your PrivateKey")
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


			patient:=	bind.NewKeyedTransactor(patient_PrivateKey)
			dataBaseOwner:=bind.NewKeyedTransactor(dataBaseOwner_PrivateKey)
			insuranceCompany:=bind.NewKeyedTransactor(insuranceCompany_PrivateKey)



			con_address_registration = common.HexToAddress("Address")
			con_address_P_HA = common.HexToAddress("Address")
			con_address_P_DBO = common.HexToAddress("Address")
			con_address_P_IC = common.HexToAddress("Address")






			var wg sync.WaitGroup


			channel_PatientIC_IC_patientId:=make(chan int)
			channel_PatientIC_IC_insuranceCompanyId:=make(chan int)
			channel_PatientIC_IC_claimId:=make(chan int)
			channel_PatientIC_DO_SendIcDetails:=make(chan []byte)
			channel_DOIC_insuranceCompany_SendEncryptedFile:=make(chan []byte)

			wg.Add(3)

			go Patient_IC(patient,client,channel_PatientIC_IC_patientId,channel_PatientIC_IC_insuranceCompanyId,channel_PatientIC_IC_claimId,channel_PatientIC_DO_SendIcDetails,dataBaseOwner.From,ctx,&wg)

			go Insurance_company(insuranceCompany,client,channel_PatientIC_IC_patientId,channel_PatientIC_IC_insuranceCompanyId,channel_PatientIC_IC_claimId,channel_DOIC_insuranceCompany_SendEncryptedFile,ctx,&wg)

			go DataBaseOwner_InsuranceCompany(dataBaseOwner,client,channel_PatientIC_DO_SendIcDetails,channel_DOIC_insuranceCompany_SendEncryptedFile,dataBaseOwner_PrivateKey,ctx,&wg)

			wg.Wait()




	}

	func DataBaseOwner_InsuranceCompany(dataBaseOwner *bind.TransactOpts,client *ethclient.Client,channel_PatientIC_DO_SendIcDetails chan []byte,channel_DOIC_insuranceCompany_SendEncryptedFile chan []byte,dbprivateKey *ecdsa.PrivateKey,ctx context.Context,wg *sync.WaitGroup){


				prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : DataBaseOwner_InsuranceCompany "



				instance_P_IC_DBO,err:=import_contract_p_ic_dbo.NewSCPICDBO(con_address_P_IC,client)

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


								_,err=bind.WaitMined(ctx,client,tx)
			          if err!=nil{
			            log.Fatal(err)
			          }


				dbIC_KeepSignature_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time DbIC_KeepSignature : %v \n\n", dbIC_KeepSignature_stop.Sub(dbIC_KeepSignature_start))
				log.SetPrefix(prefix)
				log.Println("Tx DbIC_KeepSigOnHashOfEncFile : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")




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



	func Patient_IC(patient_IC *bind.TransactOpts,client *ethclient.Client,channel_PatientIC_IC_patientId chan int,channel_PatientIC_IC_insuranceCompanyId chan int,channel_PatientIC_IC_claimId chan int,channel_PatientIC_DO_SendIcDetails chan []byte,dataBaseOwner_Address common.Address,ctx context.Context,wg *sync.WaitGroup) {


			prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : Patient_WithInsuranceCompany "

			instance_P_IC_DBO,err:=import_contract_p_ic_dbo.NewSCPICDBO(con_address_P_IC,client)


			instance_registration,err:=import_contract_p_ic_dbo.NewSystemUsersInfo(con_address_registration,client)
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
						appId:=1
						var insuranceTermsAndConditionsFileHash [32]byte

						tx,err=instance_P_IC_DBO.BuyPolicyPhaseOne(&bind.TransactOpts{
								From:patient_IC.From,
								Signer:patient_IC.Signer,
								Value:big.NewInt(int64(policyPrice)),
						},con_address_registration,big.NewInt(int64(insuranceCompanyId)),insuranceTermsAndConditionsFileHash)
						if err!=nil{
								log.Fatal(err)
						}


						_,err=bind.WaitMined(ctx,client,tx)
	          if err!=nil{
	            log.Fatal(err)
	          }

			patientIC_buyPolicyPhaseOne_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time PatientIC_buyPolicyPhaseOne : %v \n\n", patientIC_buyPolicyPhaseOne_stop.Sub(patientIC_buyPolicyPhaseOne_start))
			log.SetPrefix(prefix)
			log.Println("Tx PatientIC_buyPolicyPhaseOne : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")


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


												for {


																	header, err := client.HeaderByNumber(context.Background(), nil)
																	if err != nil {
																		 log.Fatal(err)
																	}

																	block, err := client.BlockByNumber(context.Background(), header.Number)
																	if err != nil {
																			log.Fatal(err)
																	}

																	Now:=block.Time()



																	if Now.Int64()>Time.Int64()+ timeLimiitForPolicyBuying{
																		break
																	}

																	time.Sleep(time.Second*1)


												}


												tx,err=instance_P_IC_DBO.PatientWithdrawLockedPolicyBuyMoney(&bind.TransactOpts{
															From:patient_IC.From,
															Signer:patient_IC.Signer,
												},con_address_registration,big.NewInt(int64(insuranceCompanyId)))
												if err!=nil{
												    log.Fatal(err)
												}


												_,err=bind.WaitMined(ctx,client,tx)
							          if err!=nil{
							            log.Fatal(err)
							          }



						 patientIC_UnlockingAmount_ICnotLockingMoney_stop:=time.Now()
						 log.SetPrefix(prefix)
			 			 log.Printf("Time PatientIC_UnlockingAmount_ICnotLockingMoney : %v \n\n", patientIC_UnlockingAmount_ICnotLockingMoney_stop.Sub(patientIC_UnlockingAmount_ICnotLockingMoney_start))
						 log.SetPrefix(prefix)
			 			 log.Println("Tx PatientIC_PatientWithdrawLockedPolicyBuyMoney : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")

						 wg.Done()


				}


				patientIC_GeneratingClaim_start:=time.Now()



								// fmt.Println

								rsa_privateKey,rsa_publicKey:=GenerateKeyPair(int(buffer_size*32))

								bytes_rsa_privateKey:=PrivateKeyToBytes(rsa_privateKey)
								bytes_rsa_publicKey:=PublicKeyToBytes(rsa_publicKey)

								rsa_privateKey_commit:=rsa_fnKeycommit(bytes_rsa_privateKey)


								tx,err=instance_P_IC_DBO.ClaimMoney(&bind.TransactOpts{
											From:patient_IC.From,
											Signer:patient_IC.Signer,
								},con_address_registration,con_address_P_HA,con_address_P_DBO,PolicyDetailId,big.NewInt(int64(claimAmount)),big.NewInt(int64(estimatedBillId)),bytes_rsa_publicKey,rsa_privateKey_commit,dataBaseOwner_Address,big.NewInt(int64(appId)))
								if err!=nil{
										log.Fatal(err)
								}


								_,err=bind.WaitMined(ctx,client,tx)
			          if err!=nil{
			            log.Fatal(err)
			          }


				patientIC_GeneratingClaim_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientIC_GeneratingClaim : %v \n\n", patientIC_GeneratingClaim_stop.Sub(patientIC_GeneratingClaim_start))
				log.SetPrefix(prefix)
				log.Println("Tx PatientIC_GenerateClaim : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")




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
								PatientICToDataBaseOwner_object.EstimatedBillId=estimatedBillId
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


									_,err=bind.WaitMined(ctx,client,tx)
				          if err!=nil{
				            log.Fatal(err)
				          }


				patientIC_revealSecretKey_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time PatientIC_RevealKey : %v \n\n", patientIC_revealSecretKey_stop.Sub(patientIC_revealSecretKey_start))
				log.SetPrefix(prefix)
				log.Println("Tx PatientIC_RevealSecretKey : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")



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



															for {


																				header, err := client.HeaderByNumber(context.Background(), nil)
																				if err != nil {
																					 log.Fatal(err)
																				}

																				block, err := client.BlockByNumber(context.Background(), header.Number)
																				if err != nil {
																						log.Fatal(err)
																				}


																				Now:=block.Time()


																				if Now.Int64()>TimestampKeyReveal.Int64()+ timeLimitForFinalApproval{
																					break
																				}

																				time.Sleep(time.Second*1)



															}


															tx,err=instance_P_IC_DBO.SelfApproveClaim(&bind.TransactOpts{
																		From:patient_IC.From,
																		Signer:patient_IC.Signer,
															},con_address_registration,big.NewInt(int64(claimId)),big.NewInt(int64(insuranceCompanyId)))
															if err!=nil{
																	log.Fatal(err)
															}



															_,err=bind.WaitMined(ctx,client,tx)
										          if err!=nil{
										            log.Fatal(err)
										          }



							patientIC_selfApproval_stop:=time.Now()
							log.SetPrefix(prefix)
							log.Printf("Time PatientIC_selfApproval : %v \n\n", patientIC_selfApproval_stop.Sub(patientIC_selfApproval_start))
							log.SetPrefix(prefix)
							log.Println("Tx PatientIC_SelfApproval : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")
							wg.Done()

					}

					wg.Done()

	}






	func Insurance_company(insuranceCompany *bind.TransactOpts,client *ethclient.Client,channel_PatientIC_IC_patientId chan int,channel_PatientIC_IC_insuranceCompanyId chan int,channel_PatientIC_IC_claimId chan int,channel_DOIC_insuranceCompany_SendEncryptedFile chan []byte,ctx context.Context,wg *sync.WaitGroup){

			prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : InsuranceCompany "


			var tx *types.Transaction
			var DataBaseOwnerToInsuranceCompany_object DataBaseOwnerToInsuranceCompany


			instance_registration,err:=import_contract_p_ic_dbo.NewSystemUsersInfo(con_address_registration,client)
			if err!=nil{
			    log.Fatal(err)
			}

			instance_P_HA,err:=import_contract_p_h.NewSCPHA1(con_address_P_HA,client)
			if err!=nil{
			    log.Fatal(err)
			}


			_,insuranceCompanyId,_,err:=instance_registration.GetInsuranceCompanybyAddress(nil,insuranceCompany.From)
			if err!=nil{
					log.Fatal(err)
			}

			fmt.Println("\ninsuranceCompanyId : ",insuranceCompanyId,"\n")

			channel_PatientIC_IC_insuranceCompanyId<-int(insuranceCompanyId.Int64())


			instance_P_IC_DBO,err:=import_contract_p_ic_dbo.NewSCPICDBO(con_address_P_IC,client)



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


							_,err=bind.WaitMined(ctx,client,tx)
		          if err!=nil{
		            log.Fatal(err)
		          }


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


							_,err=bind.WaitMined(ctx,client,tx)
		          if err!=nil{
		            log.Fatal(err)
		          }


			insuranceCompany_LockPolicyPhaseTwo_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time InsuranceCompany_LockPolicyPhaseTwo : %v \n\n", insuranceCompany_LockPolicyPhaseTwo_stop.Sub(insuranceCompany_LockPolicyPhaseTwo_start))
			log.SetPrefix(prefix)
			log.Println("Tx InsuranceCompany_LockPolicyPhaseTwo : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")


			insuranceCompany_waitingForPatientToSendClaimId_start:=time.Now()

								claimId:=<-channel_PatientIC_IC_claimId


			insuranceCompany_waitingForPatientToSendClaimId_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time InsuranceCompany_WaitingForPatientToSendClaimId : %v \n\n", insuranceCompany_waitingForPatientToSendClaimId_stop.Sub(insuranceCompany_waitingForPatientToSendClaimId_start))


			insuranceCompany_GettingEncFileAndVerifyingDbSignature_start:=time.Now()

								byteArrayFromDataBaseOwnerToInsuranceCompany:=<-channel_DOIC_insuranceCompany_SendEncryptedFile


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


															_,err=bind.WaitMined(ctx,client,tx)
										          if err!=nil{
										            log.Fatal(err)
										          }


										insuranceCompany_lockingApprovedMoney_stop:=time.Now()
										log.SetPrefix(prefix)
										log.Printf("Time InsuranceCompany_lockingClaimeddMoney : %v \n\n", insuranceCompany_lockingApprovedMoney_stop.Sub(insuranceCompany_lockingApprovedMoney_start))
										log.SetPrefix(prefix)
										log.Println("Tx InsuranceCompany_LockClaimedMoney : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")




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


																				for {


																									header, err := client.HeaderByNumber(context.Background(), nil)
																								 	if err != nil {
																											log.Fatal(err)
																								 	}

																								  block, err := client.BlockByNumber(context.Background(), header.Number)
																								 	if err != nil {
																										 	log.Fatal(err)
																								 	}


																								  Now:=block.Time()


																									if Now.Int64()>TimestampIcLocking.Int64()+ timeLimitForKeyReveal{
																										break
																									}

																									time.Sleep(time.Second*1)


																				}


																				tx,err=instance_P_IC_DBO.WithdrawingLockedClaimedMoney(&bind.TransactOpts{
																							From:insuranceCompany.From,
																							Signer:insuranceCompany.Signer,
																				},con_address_registration,big.NewInt(int64(claimId)),big.NewInt(int64(patientId)))
																				if err!=nil{
																						log.Fatal(err)
																				}


																				_,err=bind.WaitMined(ctx,client,tx)
															          if err!=nil{
															            log.Fatal(err)
															          }



															insuranceCompany_UnlockingAmount_patientNotRevealedKey_stop:=time.Now()
															log.SetPrefix(prefix)
															log.Printf("Time InsuranceCompany_UnlockingAmount_patientNotRevealedKey : %v \n\n", insuranceCompany_UnlockingAmount_patientNotRevealedKey_stop.Sub(insuranceCompany_UnlockingAmount_patientNotRevealedKey_start))
															log.SetPrefix(prefix)
															log.Println("Tx InsuranceCompany_LeaveAfterLockedMoney : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")


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
										estimatedBillId:=claimDetails_object.EstimatedBillId


										multiSigID,err:=instance_P_HA.EstimatedCostBillIDToMultiSigOnMedicalDataID(nil,estimatedBillId)
										if err!=nil{
												log.Fatal(err)
										}

										multiSigID_int:=multiSigID.Int64()


										MultiSigOnMedicalData_object,err:=instance_P_HA.MultiSigOnMedicalData(nil,big.NewInt(int64(multiSigID_int-1)))
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


																		_,err=bind.WaitMined(ctx,client,tx)
													          if err!=nil{
													            log.Fatal(err)
													          }


														insuranceCompany_ApproveClaimedMoney_stop:=time.Now()
														log.SetPrefix(prefix)
														log.Printf("Time InsuranceCompany_ApproveClaimedMoney_zero : %v \n\n", insuranceCompany_ApproveClaimedMoney_stop.Sub(insuranceCompany_ApproveClaimedMoney_start))
														log.SetPrefix(prefix)
														log.Println("Tx InsuranceCompany_ApproveClaimedMoney_zero : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")

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


														_,err=bind.WaitMined(ctx,client,tx)
									          if err!=nil{
									            log.Fatal(err)
									          }


										insuranceCompany_ApproveClaimedMoney_stop:=time.Now()
										log.SetPrefix(prefix)
										log.Printf("Time InsuranceCompany_ApproveClaimedMoney : %v \n\n", insuranceCompany_ApproveClaimedMoney_stop.Sub(insuranceCompany_ApproveClaimedMoney_start))
										log.SetPrefix(prefix)
										log.Println("Tx InsuranceCompany_ApproveClaimedMoney : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")



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
	          fmt.Println(err)
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
		        fmt.Println(err)
		      }
					return plainText

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
