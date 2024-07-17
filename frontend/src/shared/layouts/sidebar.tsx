import React, { useState } from 'react';
import { Outlet } from 'react-router';

import write from '@assets/icon/system-uicons_write.svg';
import tabler from '@assets/icon/tabler_layout-sidebar-filled.svg';

import LinkButton from '@shared/ui/link-button';
import { Button } from '@shared/ui';

import classes from './css/sidebar.module.css';

export default function Sidebar() {
  const [isSidebarVisible, setSidebarVisible] = useState(true);

  const handleToggleSidebar = () => {
    setSidebarVisible(!isSidebarVisible);
  };

  return (
    <div className={classes.layout}>
      <div
        className={`${classes.layout__sidebar} ${
          isSidebarVisible ? classes.visible : classes.hidden
        }`}
      >
        <header className={classes.sidebar__header}>
          <Button className={classes.header__button} onClick={handleToggleSidebar}>
            <img src={tabler} alt="사이드바 토글 버튼" />
          </Button>

          <LinkButton to="new" className={classes.header__button}>
            <img src={write} alt="새 창" />
          </LinkButton>
        </header>
      </div>

      <main className={classes.layout__chat}>
        {isSidebarVisible && <header className={classes.chat__header}></header>}
        {!isSidebarVisible && (
          <header className={classes.chat__header}>
            <Button className={classes.header__button} onClick={handleToggleSidebar}>
              <img src={tabler} alt="사이드바 토글 버튼" />
            </Button>

            <LinkButton to="new" className={classes.header__button}>
              <img src={write} alt="새 창" />
            </LinkButton>

            <h1>오늘 뭐먹지?</h1>
          </header>
        )}
        <main className={classes['chat-container']}>
          <Outlet />
        </main>
      </main>
    </div>
  );
}
