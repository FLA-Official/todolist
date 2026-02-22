import { API_BASE } from "../config/api";

export async function api(path, options = {}) {
  return fetch(`${API_BASE}${path}`, options);
}