export type MessageProps = {
  name: string;
  text: string;
};

function Message(props: MessageProps) {
  return (
    <>
      <strong>{props.name}</strong>: {props.text}
    </>
  );
}

export default Message;
