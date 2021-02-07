export const state = () => ({
    current: null,
})

export const mutations = {
    setCurrent(state, user) {
        state.current = user
    }
}

export const actions = {
    async signup(context,{nickname,password,rePassword,captchaId,captchaCode}){
        const resp=await this.$axios.post('/api/login/signup',{
            nickname,password,rePassword,captchaId,captchaCode
        })
        console.log(resp)
        return resp
    },
}

