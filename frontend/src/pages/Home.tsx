import React from 'react';
import { Layout, Typography } from 'antd';
const { Header, Content, Footer } = Layout;
const { Title } = Typography;

const Home: React.FC = () => {
  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Header style={{ color: 'white' }}>To-Do App</Header>
      <Content style={{ padding: '2rem' }}>
        <Title level={2}>Welcome to To-Do App</Title>
        {/* nanti taruh TodoList, TodoForm, dll */}
      </Content>
      <Footer style={{ textAlign: 'center' }}>© 2025 To-Do App</Footer>
    </Layout>
  );
};

export default Home;
