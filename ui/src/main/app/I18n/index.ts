// @flow
import LocalizedStrings from 'react-localization'
import _en from '../I18n/en'
import _de from '../I18n/de'

type I18nType = {
  setLanguage: (language: string) => void,
  getLanguage: () => string,
  getInterfaceLanguage: () => string,
  formatString: (str: string, ...values: any[]) => string,
  getAvailableLanguages: () => string[],
  getString: (key: string, language: string) => string,
  t: (key: string, language: string) => string,
  setContent: (props: any) => void,
  [string]: string | { [string]: string | number }
}

let I18n: I18nType = new LocalizedStrings({
  en: _en,
  de: _de
})

I18n.t = I18n.getString

export default I18n
