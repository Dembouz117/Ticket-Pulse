interface TabUnderlineProps {
  text: string;
  active: boolean;
}

const TabUnderline = ({ text, active }: TabUnderlineProps) => {
  return (
    <div
      className={`pt-10 pb-4 border-b ${
        active ? "border-pink-600" : "border-gray-600"
      } justify-center items-center flex`}
    >
      <div className="text-center text-white font-normal leading-[18px] hover:text-gray-400">
        {text}
      </div>
    </div>
  );
};

export default TabUnderline;
