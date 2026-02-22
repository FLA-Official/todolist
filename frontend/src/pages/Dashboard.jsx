import Navbar from "../components/Navbar";
import Sidebar from "../components/Sidebar";
import Board from "../components/Board";

const Dashboard = () => (
  <div className="layout">
    <Sidebar />
    <div className="main">
      <Navbar />
      <Board />
    </div>
  </div>
);

export default Dashboard;