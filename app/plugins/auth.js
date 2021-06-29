export default async function ({$auth, $axios}) {
  if (!$auth.loggedIn) {
    return
  }

  const authStrategy = $auth.strategy.name;
  if(authStrategy === 'github'){
    const token = $auth.getStrategy().token.get().substr(7);
    const user = $auth.user;

    console.log('Request Github account');
    
    try {
      const data = await $axios.$post('api/auth/github', {
        token,
        ...user
      });
      
      $auth.setStrategy('local');
      $auth.strategies.local.setUserToken(data.accessToken, data.refreshToken);
      $auth.setUser(data.user);
    } catch (e) {
      console.log(e);
    }
  }
}
