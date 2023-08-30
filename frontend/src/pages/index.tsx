import { NextPage } from "next";
import { TestService } from "src/generated/service.pb";
import useSWR from "swr";

const hello = async () => {
  const res = await TestService.Hello(
    { name: "John" },
    {
      pathPrefix: "http://localhost:8081",
    }
  );
  return res.msg;
};

const Home: NextPage = () => {
  const { data } = useSWR("hello", hello);
  return data;
};

export default Home;
