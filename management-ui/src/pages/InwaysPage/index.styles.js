// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'
import { Table } from '@commonground/design-system'
import { IconInway } from '../../icons'

export const StyledIconTd = styled(Table.Td)`
  width: 3rem;
`

export const StyledInwayIcon = styled(IconInway)`
  display: block;
  width: 1rem;
  height: 1rem;
  fill: ${(p) => p.theme.tokens.colorPaletteGray50};
`
