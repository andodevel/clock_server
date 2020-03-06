/* eslint-disable max-len */
type LanguageType = {
  [string]: {
    [string]:
      | string
      | {
          [string]: string
        }
  }
}

const en: LanguageType = {
  language: {
    en: 'Englisch',
    de: 'Deutsch'
  },
}

export default de
