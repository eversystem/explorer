
// const APIV1 = 'http://47.244.109.92:8008/api/'
const APIV1 = '/api/'

const config = {
  apis: {
    market: `${APIV1}market`,
    indexBlocks: `${APIV1}indexBlocks`,
    indexTxns: `${APIV1}indexTxns`,

    blocks: `${APIV1}blocks`,
    block: `${APIV1}block/`,

    txs: `${APIV1}txs`,
    tx: `${APIV1}tx/`,

    accounts: `${APIV1}accounts`,
    account: `${APIV1}account/`,

    contract: `${APIV1}contract/`,

    search: `${APIV1}search/`,

    sendSMS: `${APIV1}sendSMS`,

    feedback: `${APIV1}feedback`,

    applyIOST: `${APIV1}applyIOST`
  },
}

export {
  config
};