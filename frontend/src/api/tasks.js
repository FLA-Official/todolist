import { api } from "./client";

export function getTasks(projectId) {
  return api(`/projects/${projectId}/tasks`);
}

export function updateTaskStatus(taskId, status) {
  return api(`/tasks/${taskId}/status`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ status }),
  });
}