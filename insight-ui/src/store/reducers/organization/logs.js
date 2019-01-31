import cfg from '../../app.cfg'
import * as actionType from '../../actions'

function mergePageDef(state, payload) {
    let pageDef = {
        ...state.pageDef,
    }
    if (typeof payload.page !== 'undefined') {
        pageDef.page = payload.page
    }
    if (typeof payload.rowCount !== 'undefined') {
        pageDef.rowCount = payload.rowCount
    } else {
        pageDef.rowCount = payload.items.length
    }
    if (typeof payload.rowsPerPage !== 'undefined') {
        pageDef.rowsPerPage = payload.rowsPerPage
    }
    return pageDef
}

const logs = (state = cfg.organization.logs, action) => {
    switch (action.type) {
        case actionType.GET_ORGANIZATION_LOGS_OK:
            let pageDef = mergePageDef(state, action.payload)

            return {
                ...state,
                ...action.payload,
                error: null,
                pageDef: {
                    ...pageDef,
                },
            }
        case actionType.GET_ORGANIZATION_LOGS_ERR:
            return {
                ...state,
                items: [],
                error: {
                    ...action.payload,
                },
            }
        case actionType.RESET_ORGANIZATION:
        case actionType.RESET_ORGANIZATION_LOGS:
            return {
                ...cfg.organization.logs,
            }
        default:
            return state
    }
}

export default logs
