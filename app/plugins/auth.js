export default async function ({$auth}) {
  console.log($auth)
  if (!$auth.loggedIn) {
    return
  }

  const authStrategy = $auth.strategy.name;
  if(authStrategy === 'github'){
    const token = $auth.getStrategy().token.get().substr(7);
    const user = $auth.user;

    try {
      const {accessToken, refreshToken} = await app.$axios.$post('api/auth/github', {
        token,
        user
      });
      
      $auth.strategies.local.setUserToken(accessToken, refreshToken)
      setTimeout( async () => {
        auth.setStrategy('local');
      });
    } catch (e) {
      console.log(e);
    }
  }
}
