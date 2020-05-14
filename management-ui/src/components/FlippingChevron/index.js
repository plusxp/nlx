// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { bool, number } from 'prop-types'
import styled, { css } from 'styled-components'
import React from 'react'
import { ReactComponent as ChevronDown } from './chevron-down.svg'

const FlippingIconChevron = ({
  flipHorizontal,
  animationDuration,
  ...props
}) => <ChevronDown {...props} />

FlippingIconChevron.propTypes = {
  flipHorizontal: bool,
  animationDuration: number,
}
FlippingIconChevron.defaultProps = {
  animationDuration: 150,
}

const FlippingChevron = styled(FlippingIconChevron)`
  fill: ${(p) => p.theme.colorText};
  transition: ${(p) => p.animationDuration}ms ease-in-out;

  ${(p) =>
    p.flipHorizontal
      ? css`
          transform: rotate(180deg);
        `
      : ''}
`

export default FlippingChevron
