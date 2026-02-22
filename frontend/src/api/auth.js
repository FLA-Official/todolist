import { api } from "./client";

export function login(data) {
  return api("/auth/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
}

export function register(data) {
  return api("/auth/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
}

export function getProfile(token) {
  return api("/users/me", {
    headers: { Authorization: `Bearer ${token}` },
  });
}