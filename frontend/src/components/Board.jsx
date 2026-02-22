import { DragDropContext, Droppable } from "@hello-pangea/dnd";
import Column from "./Column";
import { useEffect, useState } from "react";
import { getTasks, updateTaskStatus } from "../api/tasks";

const Board = () => {
  const [columns, setColumns] = useState({
    "todo": { name: "To Do", items: [] },
    "inprogress": { name: "In Progress", items: [] },
    "done": { name: "Done", items: [] }
  });

  useEffect(() => {
    // Load tasks from backend
    getTasks("defaultProjectId") // replace with your projectId
      .then(res => res.json())
      .then(tasks => {
        const col = { todo: [], inprogress: [], done: [] };
        tasks.forEach(task => col[task.status]?.push(task));
        setColumns({
          todo: { name: "To Do", items: col.todo },
          inprogress: { name: "In Progress", items: col.inprogress },
          done: { name: "Done", items: col.done }
        });
      });
  }, []);

  const onDragEnd = async (result) => {
    const { source, destination } = result;
    if (!destination) return;

    const sourceCol = columns[source.droppableId];
    const destCol = columns[destination.droppableId];
    const [removed] = sourceCol.items.splice(source.index, 1);

    removed.status = destination.droppableId;
    destCol.items.splice(destination.index, 0, removed);

    setColumns({ ...columns });

    // Update backend
    await updateTaskStatus(removed.id, removed.status);
  };

  return (
    <DragDropContext onDragEnd={onDragEnd}>
      <div className="board">
        {Object.entries(columns).map(([id, col]) => (
          <Droppable droppableId={id} key={id}>
            {(provided) => (
              <Column
                innerRef={provided.innerRef}
                {...provided.droppableProps}
                column={col}
              >
                {provided.placeholder}
              </Column>
            )}
          </Droppable>
        ))}
      </div>
    </DragDropContext>
  );
};

export default Board;