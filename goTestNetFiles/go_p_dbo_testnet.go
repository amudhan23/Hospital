package main

import(

		"fmt"
		 kec "github.com/ethereum/go-ethereum/crypto"
		"os"
		"io"
		"log"
		"encoding/json"
		"context"
		"time"
		"sync"
		"reflect"
		"strconv"
		"crypto/rand"
	  "crypto/ecdsa"
		"github.com/ethereum/go-ethereum/ethclient"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
		"github.com/ethereum/go-ethereum/crypto"
	  "github.com/ethereum/go-ethereum/common"
		"github.com/ethereum/go-ethereum/core/types"
		"math/big"
		 mathlib "math"
	   import_contract_p_dbo "../compiledGoFiles_testNet/contract_p_dbo"
		 import_registration "../compiledGoFiles_testNet/contract_registration"

)


var buffer_size uint
var noOfInputGates uint
var totalNumberOfGates uint


var con_address_registration common.Address
var con_address_P_HA common.Address
var con_address_P_DBO common.Address


const keySize=32
const maxLinesToGate =2
const hashFunctionOutputBitSize=32



const timeLimitForMerkleRootVerification_P_DO=100
const timeLimitForKeyReveal_P_DO=100
const timeLimitForGivingComplainOrPuttingInTheDatatbase_P_DO=100


type PatientToDataBaseOwner struct {
		EncryptedOutputOfGates [][]byte
}

func main(){


			noOfInputGates=32
			buffer_size=32
			totalNumberOfGates=2*noOfInputGates

			logFilePath:="../log_testNet/log_p_dbo_testNet.txt"

			var err error
			f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
					// log.Fatalf("error opening file: %v", err)
			defer f.Close()
			}
			wrt := io.MultiWriter(os.Stdout, f)
			log.SetOutput(wrt)



			client, err := ethclient.Dial("Infura Key Point")
			if err != nil {
				log.Fatal(err)
			}

			ctx := context.Background()

			patient_PrivateKey,err:=crypto.HexToECDSA("Your PrivateKey")
			if err != nil {
				log.Fatal(err)
			}

			dataBaseOwner_PrivateKey, err := crypto.HexToECDSA("Your PrivateKey")
			if err != nil {
				log.Fatal(err)
			}


			patient:=	bind.NewKeyedTransactor(patient_PrivateKey)
			dataBaseOwner:=bind.NewKeyedTransactor(dataBaseOwner_PrivateKey)

			con_address_registration = common.HexToAddress("Address")
			con_address_P_HA = common.HexToAddress("Address")
			con_address_P_DBO = common.HexToAddress("Address")


			channel_Patient_DO_appId:=make(chan int)
			channel_Patient_DO_EncryptedFile_Object:=make(chan []byte)
			channel_Patient_DO_DBOId:=make(chan int)

			var wg sync.WaitGroup
			wg.Add(2)

			go DatabaseOwner(dataBaseOwner,client,channel_Patient_DO_appId,channel_Patient_DO_EncryptedFile_Object,channel_Patient_DO_DBOId,ctx,&wg)


			go Patient_Do(patient,client,channel_Patient_DO_appId,channel_Patient_DO_EncryptedFile_Object,channel_Patient_DO_DBOId,ctx,&wg)

			wg.Wait()


	}


	func Patient_Do(patient_DO *bind.TransactOpts,client *ethclient.Client,channel_Patient_DO_appId chan int,channel_Patient_DO_EncryptedFile_Object chan []byte,channel_Patient_DO_DBOId chan int,ctx context.Context,wg *sync.WaitGroup){

			prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : Patient_WithDataBaseOwner "

			instance_P_DO,err:=import_contract_p_dbo.NewSCPDBO(con_address_P_DBO,client)
			instance_registration,err:=import_registration.NewSystemUsersInfo(con_address_registration,client)

			var tx *types.Transaction
			var PatientToDataBaseOwner_Object PatientToDataBaseOwner

			patientId:=1
			estimatedBillId:=1


			dboId_int:=<-channel_Patient_DO_DBOId



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
							},con_address_registration,con_address_P_HA,big.NewInt(int64(estimatedBillId)),MRencout,MRinp,dboAddress)
							if err!=nil{
							    log.Fatal(err)
							}
							_,err=bind.WaitMined(ctx,client,tx)
							if err!=nil{
								log.Fatal(err)
							}


			patientDO_keepData_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time PatientDO_keepData : %v \n\n", patientDO_keepData_stop.Sub(patientDO_keepData_start))
			log.SetPrefix(prefix)
			log.Println("Tx KeepDataPatient : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")



			patientDO_sendingAppIdToDO_start:=time.Now()

							app_Id,err:=instance_P_DO.PatientIdToApplicationId(nil,big.NewInt(int64(patientId)))
							if err!=nil{
									log.Fatal(err)
							}

							fmt.Println("\n\napp_Id : ",app_Id,"\n\n")


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


										for  {


																header, err := client.HeaderByNumber(context.Background(), nil)
															 	if err != nil {
																	log.Fatal(err)
															 	}


																block, err := client.BlockByNumber(context.Background(), header.Number)
																if err != nil {
																	 log.Fatal(err)
																}


												 			  Now:=block.Time()


																if Now.Int64()>TimestampApplication.Int64()+timeLimitForMerkleRootVerification_P_DO {
																	break
																}

																time.Sleep(time.Second*1)


										}


										tx,err=instance_P_DO.LeaveAfterKeepDataPatient(&bind.TransactOpts{
												From:patient_DO.From,
												Signer:patient_DO.Signer,
										},con_address_registration,con_address_P_HA,app_Id,big.NewInt(int64(estimatedBillId)))
										if err!=nil{
										    log.Fatal(err)
										}


										_,err=bind.WaitMined(ctx,client,tx)
										if err!=nil{
											log.Fatal(err)
										}


					patientDO_UnlockingAmount_DBNotVerifiedEncFile_stop:=time.Now()
					log.SetPrefix(prefix)
					log.Printf("Time PatientDO_UnlockingAmount_DBNotVerifiedEncFile : %v \n\n", patientDO_UnlockingAmount_DBNotVerifiedEncFile_stop.Sub(patientDO_UnlockingAmount_DBNotVerifiedEncFile_start))
					log.SetPrefix(prefix)
					log.Println("Tx PatientDO_LeaveAfterKeepDataPatient : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")

					wg.Done()

			}


			patientDO_KeyReveal_start:=time.Now()

								tx,err=instance_P_DO.KeyRevealPatient(&bind.TransactOpts{
										From:patient_DO.From,
										Signer:patient_DO.Signer,
								},con_address_registration,con_address_P_HA,app_Id,key)
								if err!=nil{
										log.Fatal(err)
								}


								_,err=bind.WaitMined(ctx,client,tx)
								if err!=nil{
									log.Fatal(err)
								}


			patientDO_KeyReveal_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time PatientDO_KeyReveal : %v \n\n", patientDO_KeyReveal_stop.Sub(patientDO_KeyReveal_start))
			log.SetPrefix(prefix)
			log.Println("Tx PatientDO_KeyReveal : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")



			patientDO_WaitingForDbToGiveApprovalOrComplain_start:=time.Now()

							patientDO_ApprovalOrComplain:=false

							for i:=0;i<timeLimitForGivingComplainOrPuttingInTheDatatbase_P_DO;i++ {

								ApplicationForStoring_Object,err:=instance_P_DO.ApplicationForStoring(nil,big.NewInt(int64(app_Id.Int64()-1)))
								if err!=nil{
										log.Fatal(err)
								}

								if ApplicationForStoring_Object.TimestampComplainOrApproval.Int64()!=0	{
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


										for  {


																header, err := client.HeaderByNumber(context.Background(), nil)
																if err != nil {
																	 log.Fatal(err)
																}

																block, err := client.BlockByNumber(context.Background(), header.Number)
																if err != nil {
																		log.Fatal(err)
																}


																Now:=block.Time()


																if Now.Int64()>TimestampKeyReveal.Int64()+timeLimitForGivingComplainOrPuttingInTheDatatbase_P_DO {
																	break
																}

																time.Sleep(time.Second*1)

										}


										tx,err=instance_P_DO.LeaveAfterKeyRevealPatient(&bind.TransactOpts{
												From:patient_DO.From,
												Signer:patient_DO.Signer,
										},con_address_registration,con_address_P_HA,app_Id,big.NewInt(int64(estimatedBillId)))
										if err!=nil{
										    log.Fatal(err)
										}


										_,err=bind.WaitMined(ctx,client,tx)
										if err!=nil{
											log.Fatal(err)
										}


					patientDO_UnlockingAmount_DBNotGivenApprovalOrComplain_stop:=time.Now()
					log.SetPrefix(prefix)
					log.Printf("Time PatientDO_WaitingForDbToGiveApprovalOrComplain : %v \n\n", patientDO_UnlockingAmount_DBNotGivenApprovalOrComplain_stop.Sub(patientDO_UnlockingAmount_DBNotGivenApprovalOrComplain_start))
					log.SetPrefix(prefix)
					log.Println("Tx PatientDO_LeaveAfterKeyRevealPatient : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")
					wg.Done()

			}


			wg.Done()




	}

	func DatabaseOwner(dataBaseOwner *bind.TransactOpts,client *ethclient.Client,channel_Patient_DO_appId chan int,channel_Patient_DO_EncryptedFile_Object chan []byte,channel_Patient_DO_DBOId chan int,ctx context.Context,wg *sync.WaitGroup){

			prefix:=" buffer_size : "+strconv.FormatUint(uint64(buffer_size),10)+" noOfInputGates : "+strconv.FormatUint(uint64(noOfInputGates),10)+" party : DataBaseOwner "

			instance_P_DO,err:=import_contract_p_dbo.NewSCPDBO(con_address_P_DBO,client)

			var tx *types.Transaction
			var PatientToDataBaseOwner_Object PatientToDataBaseOwner

			instance_registration,err:=import_registration.NewSystemUsersInfo(con_address_registration,client)



			_,dboId,_,err:=instance_registration.DboByAddress(nil,dataBaseOwner.From)
			if err!=nil{
					log.Fatal(err)
			}


			channel_Patient_DO_DBOId<-int(dboId.Int64())




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


								_,err=bind.WaitMined(ctx,client,tx)
								if err!=nil{
									log.Fatal(err)
								}


			dataBaseOwner_verifyingEncFile_stop:=time.Now()
			log.SetPrefix(prefix)
			log.Printf("Time DataBaseOwner_VerifyingEncFile : %v \n\n", dataBaseOwner_verifyingEncFile_stop.Sub(dataBaseOwner_verifyingEncFile_start))
			log.SetPrefix(prefix)
			log.Println("Tx DataBaseOwnerVerifyingEncFile : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")



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



											for  {


																		header, err := client.HeaderByNumber(context.Background(), nil)
																		if err != nil {
																			 log.Fatal(err)
																		}

																		block, err := client.BlockByNumber(context.Background(), header.Number)
																		if err != nil {
																				log.Fatal(err)
																		}


																		Now:=block.Time()


																		if Now.Int64()>TimestampMerkleRootVerification.Int64()+timeLimitForKeyReveal_P_DO {
																			break
																		}


																		time.Sleep(time.Second*1)

										}


										tx,err=instance_P_DO.LeaveAfterVerifyEncFileDO(&bind.TransactOpts{
												From:dataBaseOwner.From,
												Signer:dataBaseOwner.Signer,
										},big.NewInt(int64(app_Id)))
										if err!=nil{
												log.Fatal(err)
										}


										_,err=bind.WaitMined(ctx,client,tx)
										if err!=nil{
											log.Fatal(err)
										}



						dataBaseOwner_UnlockingPatientNotRevealedKey_stop:=time.Now()
						log.SetPrefix(prefix)
						log.Printf("Time DataBaseOwner_UnlockingPatientNotRevealedKey : %v \n\n", dataBaseOwner_UnlockingPatientNotRevealedKey_stop.Sub(dataBaseOwner_UnlockingPatientNotRevealedKey_start))
						log.SetPrefix(prefix)
						log.Println("Tx DataBaseOwnerLeaveAfterVerifyEncFileDO : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")


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

								_,err=bind.WaitMined(ctx,client,tx)
								if err!=nil{
									log.Fatal(err)
								}


					dataBaseOwner_approval_stop:=time.Now()
					log.SetPrefix(prefix)
					log.Printf("Time DataBaseOwner_approval : %v \n\n", dataBaseOwner_approval_stop.Sub(dataBaseOwner_approval_start))
					log.SetPrefix(prefix)
					log.Println("Tx DataBaseOwner_approval : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")

			}



			if complain!=nil{



							dataBaseOwner_Complain_start:=time.Now()



										tx,err:=instance_P_DO.Complain(&bind.TransactOpts{
												From:dataBaseOwner.From,
												Signer:dataBaseOwner.Signer,
										},con_address_registration,con_address_P_HA,complain,encodedVectors_ThisGate,big.NewInt(int64(gateIndexes[0])),big.NewInt(int64(app_Id)))
										if err!=nil{
												log.Fatal(err)
										}


										_,err=bind.WaitMined(ctx,client,tx)
										if err!=nil{
											log.Fatal(err)
										}




						dataBaseOwner_Complain_stop:=time.Now()
						log.SetPrefix(prefix)
						log.Printf("Time DataBaseOwner_Complain : %v \n\n", dataBaseOwner_Complain_stop.Sub(dataBaseOwner_Complain_start))
						log.SetPrefix(prefix)
						log.Println("Tx DataBaseOwner_Complain : ",tx.Gas(),"\n"," txHash : ",tx.Hash().Hex(),"\n")

						wg.Done()

					}


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
