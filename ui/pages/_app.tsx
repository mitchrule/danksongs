import "../styles/globals.css";
import type { AppProps } from "next/app";
import Title from "../components/Title";
import NavBar from "../components/Navbar";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <NavBar />
    </>
  );
}
export default MyApp;
