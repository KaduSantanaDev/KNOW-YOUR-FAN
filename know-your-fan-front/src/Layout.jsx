import React from 'react';
import Nav from './components/NavComponent/Nav';
import { Outlet } from 'react-router-dom';

export default function Layout() {
  return (
    <>
      <Nav />
      <main>
        <Outlet />
      </main>
    </>
  );
}
