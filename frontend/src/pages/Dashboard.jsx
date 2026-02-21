import Sidebar from "../components/Sidebar";
import Navbar from "../components/Navbar";
import Board from "../components/Board";

const Dashboard = () => {
  return (
    <div className="app-container">
      <Sidebar />
      <div className="main-content">
        <Navbar />
        <Board />
      </div>
    </div>
  );
};

export default Dashboard;
