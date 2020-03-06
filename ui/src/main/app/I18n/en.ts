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
    en: 'English',
    de: 'German'
  },
}

export default en
