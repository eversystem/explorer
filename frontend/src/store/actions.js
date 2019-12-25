import axios from 'axios'
import base58 from 'bs58'
import types from './mutation_types'
import { config } from '../utils/config'

const { apis } = config

export default {
  async getIndexBlockList ({commit}) {
    const { data } = await axios.get(apis.indexBlocks)
    if(data.code == 0){
      commit(types.INDEXBLOCKLIST, {indexBlockList: data.data.splice(0,5)})
    }
  },

  async getIndexTxnList ({commit}) {
    const { data } = await axios.get(apis.indexTxns)
    if(data.code == 0){
      commit(types.INDEXTXNLIST, {indexTxnList: data.data.splice(0,5)})
    }
  },

  async getBlockInfo ({commit}, pages) {
    const { data } = await axios.get(`${apis.blocks}?p=${pages}`)
    if(data.code == 0){
      commit(types.BLOCKINFO, {blockInfo: data.data})
    }
  },

  async getBlockDetail ({commit}, pages) {
    const { data } = await axios.get(`${apis.block}${pages}`)
    if(data.code == 0){
      commit(types.BLOCKDETAIL, {blockDetail: data.data})
    }
  },

  async getTxnInfo ({commit}, {page, address='', blk=''}) {

    const { data } = await axios.get(`${apis.txs}?page=${page}&account=${address}&block=${blk}`)

    if(data.code == 0){
    commit(types.TXNINFO, {txnInfo: data.data})
    }
  },

  async getTxnDetail ({commit}, txHash) {
    const { data } = await axios.get(`${apis.tx}${txHash}`)
    if(data.code == 0){
      commit(types.TXNDETAIL, {txnDetail: data.data})
    }
  },

  async getAccountInfo ({commit}, pages) {
    const { data } = await axios.get(`${apis.accounts}?p=${pages}`)
    if(data.code == 0){
      commit(types.ACCOUNTINFO, {accountInfo: data.data})
    }
  },

  async getAccountDetail ({commit}, address) {
    const { data } = await axios.get(`${apis.account}${address}`)
    if(data.code == 0){
      commit(types.CONTRACTDETAIL, {contractDetail: data.data})
    }
  },

  async getContractDetail ({commit}, address) {
    const { data } = await axios.get(`${apis.account}${address}`)
    if(data.code == 0){
      commit(types.CONTRACTDETAIL, {contractDetail: data.data})
    }
  },

  async getAccountTxnInfo ({commit}, address) {
    const { data } = await axios.get(`${apis.account}${address}/txs`)
    if(data.code == 0){
      commit(types.ACCOUNTTXNINFO, {accountTxnInfo: data.data})
    }
  },
}
