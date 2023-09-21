import net from "k6/x/net";

export const options = {
  stages: [
    { duration: "2s", target: 5 },
    { duration: "3s", target: 5 },
    { duration: "5s", target: 10 },
    { duration: "5s", target: 0 },
  ],
  thresholds: {
    "http_reqs{expected_response:true}": ["rate>10"],
  },
};

export default async function() {
  const socket = await net.open("192.168.1.1:80");

  await socket.write("GET / HTTP/1.1\r\n\r\n");
}
