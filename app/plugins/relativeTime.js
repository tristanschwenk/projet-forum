export default (ctx, inject) =>{
    const relativeTime = (date) => {
        let time;
        time = ctx.$dayjs(date).fromNow();
        time = time[0].toUpperCase() + time.slice(1)
        return time
        }
    
    inject('relativeTime', relativeTime)
}


