// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgSubmitEncryptedTx } from "./types/fairyring/pep/tx";
import { MsgGetGeneralKeyshare } from "./types/fairyring/pep/tx";
import { MsgRequestGeneralKeyshare } from "./types/fairyring/pep/tx";
import { MsgCreateAggregatedKeyShare } from "./types/fairyring/pep/tx";

import { AggregatedKeyShare as typeAggregatedKeyShare} from "./types"
import { EncryptedTx as typeEncryptedTx} from "./types"
import { EncryptedTxArray as typeEncryptedTxArray} from "./types"
import { PepPacketData as typePepPacketData} from "./types"
import { NoData as typeNoData} from "./types"
import { CurrentKeysPacketData as typeCurrentKeysPacketData} from "./types"
import { CurrentKeysPacketAck as typeCurrentKeysPacketAck} from "./types"
import { Params as typeParams} from "./types"
import { TrustedCounterParty as typeTrustedCounterParty} from "./types"
import { PepNonce as typePepNonce} from "./types"
import { ActivePubKey as typeActivePubKey} from "./types"
import { QueuedPubKey as typeQueuedPubKey} from "./types"

export { MsgSubmitEncryptedTx, MsgGetGeneralKeyshare, MsgRequestGeneralKeyshare, MsgCreateAggregatedKeyShare };

type sendMsgSubmitEncryptedTxParams = {
  value: MsgSubmitEncryptedTx,
  fee?: StdFee,
  memo?: string
};

type sendMsgGetGeneralKeyshareParams = {
  value: MsgGetGeneralKeyshare,
  fee?: StdFee,
  memo?: string
};

type sendMsgRequestGeneralKeyshareParams = {
  value: MsgRequestGeneralKeyshare,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateAggregatedKeyShareParams = {
  value: MsgCreateAggregatedKeyShare,
  fee?: StdFee,
  memo?: string
};


type msgSubmitEncryptedTxParams = {
  value: MsgSubmitEncryptedTx,
};

type msgGetGeneralKeyshareParams = {
  value: MsgGetGeneralKeyshare,
};

type msgRequestGeneralKeyshareParams = {
  value: MsgRequestGeneralKeyshare,
};

type msgCreateAggregatedKeyShareParams = {
  value: MsgCreateAggregatedKeyShare,
};


export const registry = new Registry(msgTypes);

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	const structure: {fields: Field[]} = { fields: [] }
	for (let [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgSubmitEncryptedTx({ value, fee, memo }: sendMsgSubmitEncryptedTxParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSubmitEncryptedTx: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSubmitEncryptedTx({ value: MsgSubmitEncryptedTx.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSubmitEncryptedTx: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgGetGeneralKeyshare({ value, fee, memo }: sendMsgGetGeneralKeyshareParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgGetGeneralKeyshare: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgGetGeneralKeyshare({ value: MsgGetGeneralKeyshare.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgGetGeneralKeyshare: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRequestGeneralKeyshare({ value, fee, memo }: sendMsgRequestGeneralKeyshareParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRequestGeneralKeyshare: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRequestGeneralKeyshare({ value: MsgRequestGeneralKeyshare.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRequestGeneralKeyshare: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateAggregatedKeyShare({ value, fee, memo }: sendMsgCreateAggregatedKeyShareParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateAggregatedKeyShare: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateAggregatedKeyShare({ value: MsgCreateAggregatedKeyShare.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateAggregatedKeyShare: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgSubmitEncryptedTx({ value }: msgSubmitEncryptedTxParams): EncodeObject {
			try {
				return { typeUrl: "/fairyring.pep.MsgCreateAggregatedKeyShare", value: MsgCreateAggregatedKeyShare.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateAggregatedKeyShare: Could not create message: ' + e.message)
			}
		},
		
		msgGetGeneralKeyshare({ value }: msgGetGeneralKeyshareParams): EncodeObject {
			try {
				return { typeUrl: "/fairyring.pep.MsgGetGeneralKeyshare", value: MsgGetGeneralKeyshare.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgGetGeneralKeyshare: Could not create message: ' + e.message)
			}
		},
		
		msgRequestGeneralKeyshare({ value }: msgRequestGeneralKeyshareParams): EncodeObject {
			try {
				return { typeUrl: "/fairyring.pep.MsgRequestGeneralKeyshare", value: MsgRequestGeneralKeyshare.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRequestGeneralKeyshare: Could not create message: ' + e.message)
			}
		},
		
		msgCreateAggregatedKeyShare({ value }: msgCreateAggregatedKeyShareParams): EncodeObject {
			try {
				return { typeUrl: "/fairyring.pep.MsgCreateAggregatedKeyShare", value: MsgCreateAggregatedKeyShare.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateAggregatedKeyShare: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						AggregatedKeyShare: getStructure(typeAggregatedKeyShare.fromPartial({})),
						EncryptedTx: getStructure(typeEncryptedTx.fromPartial({})),
						EncryptedTxArray: getStructure(typeEncryptedTxArray.fromPartial({})),
						PepPacketData: getStructure(typePepPacketData.fromPartial({})),
						NoData: getStructure(typeNoData.fromPartial({})),
						CurrentKeysPacketData: getStructure(typeCurrentKeysPacketData.fromPartial({})),
						CurrentKeysPacketAck: getStructure(typeCurrentKeysPacketAck.fromPartial({})),
						Params: getStructure(typeParams.fromPartial({})),
						TrustedCounterParty: getStructure(typeTrustedCounterParty.fromPartial({})),
						PepNonce: getStructure(typePepNonce.fromPartial({})),
						ActivePubKey: getStructure(typeActivePubKey.fromPartial({})),
						QueuedPubKey: getStructure(typeQueuedPubKey.fromPartial({})),
						
		};
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			FairyringPep: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;