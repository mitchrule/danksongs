import "../styles/globals.css";
import type { AppProps } from "next/app";
import Title from "../components/Title";
import NavBar from "../components/Navbar";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <head>
        <style>
          @import
          url('https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap');
        </style>
      </head>
      <NavBar />
    </>
  );
}
export default MyApp;
