export const fetchCycleTimeByProject = async (projectKey) => {
    try {
        const res = await fetch(`http://localhost:8080/api/v1/cycle-time/${projectKey}`)
        const data = await res.json()
        return data
    } catch (error) {
        throw new Error(error)
    }
}

export const fetchThroughputByProject = async (projectKey) =>{
    try{
        const res = await fetch (`http://localhost:8080/api/v1/throughput/${projectKey}`)
        const data = res.json()
        return data
    } catch (error){
        throw new Error(error)
    }
}