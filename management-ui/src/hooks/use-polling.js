// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useEffect } from 'react'

export const INTERVAL = 3000

const functions = []
let timer

const addFunction = (fn) => {
  if (functions.includes(fn)) return

  functions.push(fn)

  if (functions.length < 2) {
    start()
  }
}

const removeFunction = (fn) => {
  const id = functions.indexOf(fn)
  if (id < 0) return

  functions.splice(id, 1)

  if (functions.length < 1) {
    stop()
  }
}

const start = () => {
  timer = setTimeout(() => {
    functions.forEach((fn) => fn())
    start()
  }, INTERVAL)
}

const stop = () => {
  clearTimeout(timer)
}

// Repeatedly call functions with a global timer
const usePolling = (fn) => {
  useEffect(() => {
    addFunction(fn)

    return () => {
      removeFunction(fn)
    }
  }, [fn])

  const pausePolling = () => removeFunction(fn)
  const continuePolling = () => addFunction(fn)
  return [pausePolling, continuePolling]
}

export default usePolling
