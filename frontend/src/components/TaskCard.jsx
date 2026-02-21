import { Draggable } from "@hello-pangea/dnd";

const TaskCard = ({ task, index }) => {
  return (
    <Draggable draggableId={task.id} index={index}>
      {(provided) => (
        <div
          className="task-card"
          ref={provided.innerRef}
          {...provided.draggableProps}
          {...provided.dragHandleProps}
        >
          <p>{task.title}</p>
        </div>
      )}
    </Draggable>
  );
};

export default TaskCard;
