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
		"context"
	  "crypto/ecdsa"
	 	"github.com/ethereum/go-ethereum/ethclient"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
		"github.com/ethereum/go-ethereum/crypto"
		// "github.com/ethereum/go-ethereum/core"
	  "github.com/ethereum/go-ethereum/common"
		"math/big"
		// import_registration "../compiledGoFiles_testNet/contract_registration"
		import_contract_rc_dbo "../compiledGoFiles_testNet/contract_rc_dbo"
)



const hashFunctionOutputBitSize=32


var con_address_registration common.Address
var con_address_RC_DBO common.Address
var dataBaseOwner_address common.Address



type DBToResearcCommunity struct {
		AggregatedData []byte
}



func main(){


			logFilePath:="../log_testNet/log_rc_dbo_testNet.txt"

			var err error
			f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
					log.Fatal("error opening file: %v", err)
			}
			defer f.Close()
			wrt := io.MultiWriter(os.Stdout, f)
			log.SetOutput(wrt)



			client, err := ethclient.Dial("Infura Key Point")
			if err != nil {
				log.Fatal(err)
			}

			ctx := context.Background()


			researchCommunity_privateKey, err := crypto.HexToECDSA("Your PrivateKey")
			if err != nil {
				log.Fatal(err)
			}


			dataBaseOwner_privateKey,err:=crypto.HexToECDSA("Your PrivateKey")
			if err != nil {
				log.Fatal(err)
			}



			researchCommunity:=bind.NewKeyedTransactor(researchCommunity_privateKey)
			dataBaseOwner:=bind.NewKeyedTransactor(dataBaseOwner_privateKey)
			// govtAuthority:=bind.NewKeyedTransactor(govtAuthority_privateKey)

			dataBaseOwner_address=dataBaseOwner.From


			con_address_registration = common.HexToAddress("Address")
			con_address_RC_DBO = common.HexToAddress("Address")


			var wg sync.WaitGroup

			channel_res_db_researchDataId:=make(chan int)
			channel_res_db_researchData:=make(chan []byte)


			wg.Add(2)

			go DataBaseOwner_withResearchCommunity(dataBaseOwner,dataBaseOwner_privateKey,client,channel_res_db_researchDataId,channel_res_db_researchData,ctx,&wg)

			go ResearchCommunity(researchCommunity,client,channel_res_db_researchDataId,channel_res_db_researchData,ctx,&wg)

			wg.Wait()

	}



	func DataBaseOwner_withResearchCommunity(dataBaseOwner *bind.TransactOpts,dataBaseOwnerPrivateKey *ecdsa.PrivateKey,client *ethclient.Client,channel_res_db_researchDataId chan int,channel_res_db_researchData chan []byte,ctx context.Context,wg *sync.WaitGroup){


				prefix:=" party : DataBaseOwner_withResearchCommunity "

				var DBToResearcCommunity_object DBToResearcCommunity
				var fileSize uint
				fileSize=64


				instance_research_db,err:=import_contract_rc_dbo.NewSCRCDBO(con_address_RC_DBO,client)
				if err!=nil{
						log.Fatal(err)
				}

				researchDataId_int:=<-channel_res_db_researchDataId
				file:=FileRead(fileSize)
				MR_file:=Hash(file)


				byte_MR_file:=make([]byte,hashFunctionOutputBitSize)

				for i:=0;i<hashFunctionOutputBitSize;i++{

								byte_MR_file[i]=MR_file[i]
				}


				signature_byte_MR_file, err := crypto.Sign(byte_MR_file,dataBaseOwnerPrivateKey)
				if err != nil {
					log.Fatal(err)
				}

				dataBaseOwnerResCom_giveData_start:=time.Now()

									tx,err:=instance_research_db.ProvideMedicalData(&bind.TransactOpts{
															From:dataBaseOwner.From,
															Signer:dataBaseOwner.Signer,
									},con_address_registration,big.NewInt(int64(researchDataId_int)),MR_file,signature_byte_MR_file)
									if err!=nil{
											log.Fatal(err)
									}


									_,err=bind.WaitMined(ctx,client,tx)
									if err!=nil{
										log.Fatal(err)
									}


				dataBaseOwnerResCom_giveData_stop:=time.Now()
				log.SetPrefix(prefix)
				log.Printf("Time DataBaseOwnerResCom_GiveData : %v \n\n", dataBaseOwnerResCom_giveData_stop.Sub(dataBaseOwnerResCom_giveData_start))
				log.SetPrefix(prefix)
				log.Println("Tx DataBaseOwnerResCom_GiveData : ",tx.Gas()," txHash : ",tx.Hash().Hex(),"\n")


				DBToResearcCommunity_object.AggregatedData=file
				byteArrayFromDbToResearchCommunity,err := json.Marshal(DBToResearcCommunity_object)
				if err != nil {
							log.Fatal(err)
				}

				channel_res_db_researchData<-byteArrayFromDbToResearchCommunity

				wg.Done()


	}

	func ResearchCommunity(researchCommunity *bind.TransactOpts,client *ethclient.Client,channel_res_db_researchDataId chan int,channel_res_db_researchData chan []byte,ctx context.Context,wg *sync.WaitGroup){

						prefix:=" party : ResearchCommunity "

						var DBToResearcCommunity_object DBToResearcCommunity

						instance_research_db,err:=import_contract_rc_dbo.NewSCRCDBO(con_address_RC_DBO,client)
						if err!=nil{
								log.Fatal(err)
						}



						researchCommunity_requestData_start:=time.Now()

											tx,err:=instance_research_db.RequestMedicalData(&bind.TransactOpts{
																	From:researchCommunity.From,
																	Signer:researchCommunity.Signer,
											},con_address_registration,dataBaseOwner_address)
											if err!=nil{
													log.Fatal(err)
											}

											_,err=bind.WaitMined(ctx,client,tx)
											if err!=nil{
												log.Fatal(err)
											}


						researchCommunity_requestData_stop:=time.Now()
						log.SetPrefix(prefix)
						log.Printf("Time ResearchCommunity_requestDatat : %v \n\n", researchCommunity_requestData_stop.Sub(researchCommunity_requestData_start))
						log.SetPrefix(prefix)
						log.Println("Tx ResearchCommunity_requestDatat : ",tx.Gas()," txHash : ",tx.Hash().Hex(),"\n")




						researchDataId,err:=instance_research_db.ResearchCommunityDatabaseOwnerResearchDataId(nil,researchCommunity.From,dataBaseOwner_address)
						if err!=nil{
								log.Fatal(err)
						}

						researchDataId_int:=researchDataId.Int64()

						channel_res_db_researchDataId<-int(researchDataId_int)

						byteArrayFromDbToResearchCommunity:=<-channel_res_db_researchData


						err=json.Unmarshal(byteArrayFromDbToResearchCommunity, &DBToResearcCommunity_object)
						if err != nil {
									log.Fatal(err)
						}


						researchCommunity_verifyingFile_start:=time.Now()

									file:=DBToResearcCommunity_object.AggregatedData
									MR_file:=Hash(file)

									verify_MR_file:=false
									verify_sig_MR_file:=false

									ResearchDataStruct_object,err:=instance_research_db.ResearchDataDetails(nil,big.NewInt(int64(researchDataId_int-1)))
									if err!=nil{
											log.Fatal(err)
									}


									if ResearchDataStruct_object.HashOfMedData ==MR_file{
											verify_MR_file=true
									}

									verify_sig_MR_file=VerifySignatue(dataBaseOwner_address,MR_file,ResearchDataStruct_object.SigOnHashOfMedData)

						researchCommunity_verifyingFile_stop:=time.Now()
						log.SetPrefix(prefix)
						log.Printf("Time ResearchCommunity_VerifyingFile : %v \n\n", researchCommunity_verifyingFile_stop.Sub(researchCommunity_verifyingFile_start))

						fmt.Println("\n verify_MR_file :",verify_MR_file,"\n verify_sig_MR_file : ",verify_sig_MR_file)

						wg.Done()
	}




	func FileRead(_bufferSize uint) ([]byte){


					var currentByte int64 = 0
					fileBuffer := make([]byte, _bufferSize)
					file, err := os.Open("input.txt")

							n,err := file.ReadAt(fileBuffer, currentByte)
							if err!=nil{
								log.Println(err)
							}

							if err == io.EOF {
									log.Fatal(err)
							}

							fileBuffer=fileBuffer[:n]


					file.Close()
					return fileBuffer

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
