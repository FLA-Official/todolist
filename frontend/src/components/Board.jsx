import { useState } from "react";
import { DragDropContext } from "@hello-pangea/dnd";
import Column from "./Column";

const initialTasks = [
  { id: "1", title: "Design login page", status: "todo" },
  { id: "2", title: "Fix API bug", status: "inprogress" },
  { id: "3", title: "Deploy backend", status: "done" },
];

const Board = () => {
  const [tasks, setTasks] = useState(initialTasks);

  const handleDragEnd = (result) => {
    if (!result.destination) return;

    const { draggableId, destination } = result;

    const updatedTasks = tasks.map((task) =>
      task.id === draggableId
        ? { ...task, status: destination.droppableId }
        : task
    );

    setTasks(updatedTasks);
  };

  return (
    <DragDropContext onDragEnd={handleDragEnd}>
      <div className="board">
        <Column title="To Do" status="todo" tasks={tasks} />
        <Column title="In Progress" status="inprogress" tasks={tasks} />
        <Column title="Done" status="done" tasks={tasks} />
      </div>
    </DragDropContext>
  );
};

export default Board;
