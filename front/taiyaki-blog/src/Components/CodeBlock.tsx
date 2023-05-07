import { CodeComponent } from 'react-markdown/lib/ast-to-react';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { a11yDark } from 'react-syntax-highlighter/dist/cjs/styles/prism';
import styled from 'styled-components';

const CodeBlockWrapper = styled.div`
  position: relative;
`;

const CodeBlockTitle = styled.div`
  display: inline-block;
  position: absolute;
  top: -0.5em;
  left: 0;
  background-color: #ccc;
  padding: 0.2em;
  color: #000;
`;

const CodeBlock: CodeComponent = ({ inline, className, children }) => {
  if (inline) {
    return <code className={className}>{children}</code>;
  }
  const match = /language-(\w+)/.exec(className || '');
  const lang = match && match[1] ? match[1] : '';
  const name = match && match[2] ? match[2].slice(1) : '';
  return (
    <CodeBlockWrapper>
      {name ? <CodeBlockTitle>{name}</CodeBlockTitle> : ''}
      <SyntaxHighlighter
        style={a11yDark}
        language={lang}
        children={String(children).replace(/\n$/, '')}
      />
    </CodeBlockWrapper>
  );
};

export default CodeBlock;
