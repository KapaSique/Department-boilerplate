const API_BASE = '/api/v1'

async function fetchAPI(endpoint) {
  const response = await fetch(`${API_BASE}${endpoint}`)
  if (!response.ok) {
    throw new Error(`API error: ${response.status}`)
  }
  return response.json()
}

export const api = {
  getNews: (limit = 20) => fetchAPI(`/news?limit=${limit}`),
  getNewsById: (id) => fetchAPI(`/news/${id}`),
  getDocuments: (category = '') => {
    const query = category ? `?category=${category}` : ''
    return fetchAPI(`/documents${query}`)
  },
  getDocumentById: (id) => fetchAPI(`/documents/${id}`),
  getBudgetData: (year = '') => {
    const query = year ? `?year=${year}` : ''
    return fetchAPI(`/budget${query}`)
  },
  getBudgetReports: () => fetchAPI('/budget/reports'),
  getDepartmentInfo: () => fetchAPI('/department/info'),
  getContacts: () => fetchAPI('/department/contacts'),
  healthCheck: () => fetchAPI('/health')
}
