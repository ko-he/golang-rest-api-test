import { request } from "./apiClient"

export const createCoffee = async (): Promise<string> => {
	const response = await request({
		url: '/api/v1/coffees',
		method: 'post',
		data: {
			name: "",
			comment: "",
			productionArea: "",
			richScore: 0,
			bitternessScore: 0,
			sournessScore: 0,
			aromaScore: 0
		}
	})

	return response.data.id
}

export const getCoffee = async (): Promise<string> => {
	const response = await request({
		url: '/api/v1/coffees/{id}',
		method: 'get',
		pathParam: { id: "dummy-coffee-uuid" }
	})

	return response.data.coffee.id
}
