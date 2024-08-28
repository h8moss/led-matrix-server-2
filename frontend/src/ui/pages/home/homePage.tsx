import LedMatrix from "../../services/led-matrix";

const quickActions = [
  {
    name: "Colors",
    className: "bg-green-400 hover:bg-green-500",
    request: "" // TODO
  },
  {
    name: "Game of life",
    className: "bg-blue-400 hover:bg-blue-500",
    request: "" // TODO
  },
  {
    name: "Time date",
    className: "bg-purple-400 hover:bg-purple-500",
    request: "" // TODO
  },
  {
    name: "Images",
    className: "bg-red-400 hover:bg-red-500",
    request: "" // TODO
  },
  {
    name: "Idle",
    className: "bg-zinc-400 hover:bg-zinc-500",
    request: "" // TODO
  }
]

const HomePage = () => {
  return (
    <>
      <h1>Welcome to the Led matrix server!</h1>
      <h2 className="text-xl p-8 border-b-2 border-sky-900">Quick Actions</h2>
      <div className="flex flex-wrap justify-around">
        {quickActions.map((action) => (
          <button
            className={`
              transition-all
              ${action.className} text-white font-bold 
              hover:scale-105
              py-2 p-5 m-2 
              flex-1 min-h-40 
              aspect-[5/3]
              flex flex-col-reverse
              text-2xl
              rounded-md`}
          >
            {action.name}
          </button>
        ))}
      </div>
    </>
  );
};

export default HomePage;
