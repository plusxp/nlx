// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

import React from 'react'
import { func, string } from 'prop-types'
import { StyledServiceDetailPane, StyledHeader, StyledTitle, StyledSecondTitle, StyledSubtitle, StyledDl, StyledEmailAddressLink, StyledCloseButton } from './index.styles'
import CloseIcon from '../CloseIcon'

const ServiceDetailPane = ({ name, organizationName, contactEmailAddress, closeHandler }) =>
  <StyledServiceDetailPane>
    <StyledHeader>
      <StyledTitle>
        {name}
      </StyledTitle>
      <StyledCloseButton onClick={() => closeHandler()}>
        <CloseIcon />
      </StyledCloseButton>
    </StyledHeader>

    <StyledSecondTitle>
      {organizationName}
    </StyledSecondTitle>

    {
      contactEmailAddress ?
        (
          <>
            <StyledSubtitle>Support</StyledSubtitle>
            <StyledDl>
              <dt>Email address</dt>
              <dd><StyledEmailAddressLink href={'mailto:' + contactEmailAddress  }>{contactEmailAddress}</StyledEmailAddressLink></dd>
            </StyledDl>
          </>
      ) : null
    }
  </StyledServiceDetailPane>

ServiceDetailPane.propTypes = {
  name: string,
  organizationName: string,
  contactEmailAddress: string,
  closeHandler: func,
}

ServiceDetailPane.defaultProps = {
  closeHandler: () => {}
}

export default ServiceDetailPane

