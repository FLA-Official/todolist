import { Draggable } from "@hello-pangea/dnd";
import TaskCard from "./TaskCard";

const Column = ({ column, innerRef, children, ...props }) => {
  return (
    <div className="column" ref={innerRef} {...props}>
      <h3>{column.name}</h3>
      {column.items.map((task, index) => (
        <Draggable key={task.id} draggableId={String(task.id)} index={index}>
          {(provided) => (
            <TaskCard
              task={task}
              innerRef={provided.innerRef}
              {...provided.draggableProps}
              {...provided.dragHandleProps}
            />
          )}
        </Draggable>
      ))}
      {children}
    </div>
  );
};

export default Column;