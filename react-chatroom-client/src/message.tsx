export type MessageProps = {
  name: string;
  text: string;
};

function stringToHslColor(str: string, s: number, l: number) {
  let hash = 0;
  for (let i = 0; i < str.length; i++) {
    hash = str.charCodeAt(i) + ((hash << 5) - hash);
  }

  const h = hash % 360;
  return `hsl(${h}, ${s}%, ${l}%)`;
}

function Message(props: MessageProps) {
  const nameColor = stringToHslColor(props.name, 70, 50);

  return (
    <>
      <strong style={{ color: nameColor }}>{props.name}</strong>: {props.text}
    </>
  );
}

export default Message;
