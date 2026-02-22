const BASE_URL = "http://localhost:8080";

export const api = {
  register: (data) =>
    fetch(`${BASE_URL}/auth/register`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    }),

  login: (data) =>
    fetch(`${BASE_URL}/auth/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    }),

  getProfile: (token) =>
    fetch(`${BASE_URL}/users/me`, {
      headers: { Authorization: `Bearer ${token}` },
    }),

  getTasks: (token, projectId) =>
    fetch(`${BASE_URL}/projects/${projectId}/tasks`, {
      headers: { Authorization: `Bearer ${token}` },
    }),

  updateTaskStatus: (token, taskId, status) =>
    fetch(`${BASE_URL}/tasks/${taskId}/status`, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ status }),
    }),
};