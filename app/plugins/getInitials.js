export default (_, inject) => {
  inject('getInitials', (str="") => {
    let words = str.split(' ')

    if (words.length == 0) return '';
    if (words.length == 1) return words[0].slice(0,2).toUpperCase();
    return (words[0].slice(0,1)+words[1].slice(0,1)).toUpperCase();  
  })
}
