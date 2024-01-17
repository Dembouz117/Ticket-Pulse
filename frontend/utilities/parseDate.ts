export const parseDate = (timestamp: number): string => {
    const daysOfWeek = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
    const months = [
      'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
      'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec',
    ];
  
    const date = new Date(timestamp * 1000); // Convert to milliseconds
    const dayOfWeek = daysOfWeek[date.getUTCDay()];
    const month = months[date.getUTCMonth()];
    const day = date.getUTCDate();
    const year = date.getUTCFullYear().toString().substr(-2);
    const hours = date.getUTCHours().toString().padStart(2, '0');
    const minutes = date.getUTCMinutes().toString().padStart(2, '0');
  
    const formattedDate = `${dayOfWeek}, ${month} ${day}/${year}`;
    const formattedTime = `${hours}:${minutes}`;
  
    return `${formattedDate}, ${formattedTime}`;
  }
  
  const timestamp = 1697076000;
  const formattedDateTime = parseDate(timestamp);
  console.log(formattedDateTime); // Example output: "Tue, Feb 08/22, 13:20"
  