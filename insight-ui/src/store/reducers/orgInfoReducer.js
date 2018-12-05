import cfg from '../app.cfg'
import * as actionType from '../actions'

/**
 * Language reducer containing lang info and translations
 * @param {object} state previous state object
 * @param {object} action redux action with type and payload
 * @param {string} action.type redux action type, string constant
 * @param {object} action.payload redux action type, string constant
 */

export const orgInfoReducer = (state = cfg.organization.info, action) => {
    switch (action.type) {
        case actionType.SELECT_ORGANIZATION:
            return {
                ...state,
                ...action.payload,
            }
        case actionType.RESET_ORGANIZATION:
        case actionType.RESET_ORGANIZATION_LOGS:
            return {
                ...cfg.organization.info,
            }

        default:
            return state
    }
}

export default orgInfoReducer
