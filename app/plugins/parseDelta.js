import {QuillDeltaToHtmlConverter} from 'quill-delta-to-html';

export default (_, inject) => {

  const defaultConfig = {

  }

  inject('parseDelta', (delta, cfg = defaultConfig) => {
    if(typeof delta == 'string') {
      try {
        delta = JSON.parse(delta);
      } catch (error) {
        console.log('[parseDelta]', delta, error);
        return delta;
      }
    }

    if(typeof delta === 'object' && !Array.isArray(delta) && 'ops' in delta) {
      delta = delta.ops
    }

    var converter = new QuillDeltaToHtmlConverter(delta, cfg);
    console.log('\t', converter.convert());
    return converter.convert();
  })
}
