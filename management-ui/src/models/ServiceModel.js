// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { action, decorate, flow, observable } from 'mobx'
import { arrayOf, bool, func, shape, string } from 'prop-types'
import {
  createModelSchema,
  createSimpleSchema,
  list,
  object,
  primitive,
  serialize,
} from 'serializr'

export const serviceModelPropTypes = {
  name: string.isRequired,
  endpointURL: string.isRequired,
  documentationURL: string.isRequired,
  apiSpecificationURL: string.isRequired,
  internal: bool.isRequired,
  techSupportContact: string.isRequired,
  publicSupportContact: string.isRequired,
  authorizationSettings: shape({
    mode: string.isRequired,
    authorizations: arrayOf({
      organizationName: string.isRequired,
      publicKeyHash: string.isRequired,
    }),
  }),
  inways: arrayOf(string),
  fetch: func.isRequired,
  update: func.isRequired,
}

// TODO test
class ServiceModel {
  isReady = false

  name = ''
  endpointURL = ''
  documentationURL = ''
  apiSpecificationURL = ''
  internal = false
  techSupportContact = ''
  publicSupportContact = ''
  authorizationSettings = { authorizations: [] }
  inways = []

  constructor({ store, service }) {
    this.store = store

    this.name = service.name
    this.endpointURL = service.endpointURL
    this.documentationURL = service.documentationURL
    this.apiSpecificationURL = service.apiSpecificationURL
    this.internal = service.internal
    this.techSupportContact = service.techSupportContact
    this.publicSupportContact = service.publicSupportContact
    this.authorizationSettings = service.authorizationSettings
    this.inways = service.inways

    this.isReady = true
  }

  fetch = flow(function* fetch() {
    this.isReady = false

    try {
      const service = yield this.store.domain.getByName(this.name)

      this.with(service)
    } finally {
      this.isReady = true
    }
  })

  with = function (service) {
    this.endpointURL = service.endpointURL
    this.documentationURL = service.documentationURL
    this.apiSpecificationURL = service.apiSpecificationURL
    this.internal = service.internal
    this.techSupportContact = service.techSupportContact
    this.publicSupportContact = service.publicSupportContact
    this.authorizationSettings = service.authorizationSettings
    this.inways = service.inways
  }

  update = flow(function* update(values) {
    this.isReady = false
    this.with(values)
    try {
      yield this.store.domain.update(this.name, serialize(this))

      return this
    } finally {
      this.isReady = true
    }
  })
}

createModelSchema(ServiceModel, {
  name: primitive(),
  endpointURL: primitive(),
  documentationURL: primitive(),
  apiSpecificationURL: primitive(),
  internal: primitive(),
  techSupportContact: primitive(),
  publicSupportContact: primitive(),
  authorizationSettings: object(
    createSimpleSchema({
      mode: primitive(),
      authorizations: list(
        object(
          createSimpleSchema({
            organizationName: primitive(),
            publicKeyHash: primitive(),
          }),
        ),
      ),
    }),
  ),
  inways: list(primitive()),
})

decorate(ServiceModel, {
  name: observable,
  endpointURL: observable,
  documentationURL: observable,
  apiSpecificationURL: observable,
  internal: observable,
  techSupportContact: observable,
  publicSupportContact: observable,
  authorizationSettings: observable,
  inways: observable,
  isReady: observable,
  fetch: action.bound,
  update: action.bound,
})

export const createService = (...args) => new ServiceModel(...args)

export default ServiceModel