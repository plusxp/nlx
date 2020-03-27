// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components/macro'
import FlippingChevron from '../FlippingChevron'

export const UserNavigationChevron = styled(FlippingChevron)`
  width: ${(p) => p.theme.tokens.spacing06};
  height: ${(p) => p.theme.tokens.spacing06};
  margin-left: ${(p) => p.theme.tokens.spacing04};
`

export const StyledAvatar = styled.figure`
  padding: 4px 0;
  margin: 0;
  height: ${(p) => p.theme.tokens.spacing09};

  .avatar-image {
    max-height: 100%;
    max-width: 100%;
    border-radius: 100%;
  }
`

export const StyledUserNavigation = styled.div`
  display: flex;
  align-items: center;
  position: relative;

  img {
    margin-right: ${(p) => p.theme.tokens.spacing04};
  }

  ul {
    display: block;
    position: absolute;
    top: 35px;
    right: 0;
    padding: 0;
    background: ${(p) => p.theme.tokens.colorPaletteGray800};
    box-shadow: 0 5px 20px 0 rgba(0, 0, 0, 0.25);
    list-style-type: none;
    min-width: 200px;
    z-index: 2;

    li {
      &:not(:last-child) {
        border-bottom: 1px solid #e6e6e6;
      }

      a,
      button {
        display: block;
        text-align: left;
        text-decoration: none;
        color: ${(p) => p.theme.tokens.colorText};
        font-size: ${(p) => p.theme.tokens.fontSizeMedium};
        padding: ${(p) => p.theme.tokens.spacing04}
          ${(p) => p.theme.tokens.spacing06};
        border: 2px solid transparent;
      }

      a:hover,
      button:hover,
      li:hover {
        background-color: ${(p) => p.theme.colorBackgroundDropdownHover};
      }

      a:active,
      button:active {
        background-color: ${(p) => p.theme.colorBackgroundDropdownActive};
      }

      a:focus,
      button:focus {
        border-color: ${(p) => p.theme.colorBorderDropdownFocus};
      }
    }

    button {
      width: 100%;
      padding: 0;
      border: none;
      font: inherit;
      color: inherit;
      background-color: transparent;
      cursor: pointer;
    }
  }
`

export const StyledToggleButton = styled.button`
  padding: 0;
  font: inherit;
  color: inherit;
  background-color: transparent;
  cursor: pointer;
  display: flex;
  align-items: center;
`

export const StyledUsername = styled.span`
  white-space: nowrap;
`
