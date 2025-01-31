import { request } from "./apiClient"

export const createCoffee = async (
			name: string,
			comment: string,
			productionArea: string,
			richScore: number,
			bitternessScore: number,
			sournessScore: number,
			aromaScore: number,
): Promise<string> => {
	const response = await request({
		url: '/api/v1/coffees',
		method: 'post',
		data: {
			name,
			comment,
			productionArea,
			richScore,
			bitternessScore,
			sournessScore,
			aromaScore,
		}
	})

	return response.data.id
}

export const getCoffee = async (id: string): Promise<string> => {
	const response = await request({
		url: '/api/v1/coffees/{id}',
		method: 'get',
		pathParam: { id }
	})

	return response.data.coffee.id
}
