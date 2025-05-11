package cypher

import (
	"os"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mdarin/cypher/parser"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/ffmt.v1"
)

func TestExample(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want *parser.CypherParser
	}{
		{
			name: `MATCH (:Person {name: "Alice"})-[:KNOWS]->(:Person {name: "Bob"}) RETURN count(*)`,
			args: args{query: `MATCH (:Person {name: "Alice"})-[:KNOWS]->(:Person {name: "Bob"}) RETURN count(*)`},
			want: nil,
		},
		{
			name: `CREATE (n:Person {name: 'Andy', title: 'Developer'})`,
			args: args{query: `CREATE (n:Person {name: 'Andy', title: 'Developer'})`},
			want: nil,
		},
		{
			name: `MATCH (p1:Person {name: "Alice", IsCool: $isCool})-[:KNOWS]->(p2:Person) WHERE p2.name != "Tom" RETURN count(*)`,
			args: args{query: `MATCH (p1:Person {name: "Alice", IsCool: $isCool})-[:KNOWS]->(p2:Person) WHERE p2.name != "Tom" RETURN count(*)`},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the input
			is := antlr.NewInputStream(tt.args.query)

			// Create the Lexer
			lexer := parser.NewCypherLexer(is)
			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

			// Create the Parser
			p := parser.NewCypherParser(stream)

			tw := tablewriter.NewWriter(os.Stderr)
			tw.SetColWidth(60)
			tw.SetHeader([]string{"Type", "Text"})

			// Finally parse the expression (by walking the tree)
			// The "listener" is the action triggered when visiting each node
			listener := &cypherListener{tw: tw}
			antlr.ParseTreeWalkerDefault.Walk(listener, p.OC_Cypher())

			tw.Render()
		})
	}
}

type myListener struct{}

func (x *myListener) VisitTerminal(node antlr.TerminalNode) {
	ffmt.Printf("visit terminal| %s\n", node.GetText())
}
func (x *myListener) VisitErrorNode(node antlr.ErrorNode) {
	ffmt.Printf("visit err node| %s\n", node.GetText())
}
func (x *myListener) EnterEveryRule(ruleCtx antlr.ParserRuleContext) {
	ffmt.Printf("enter every rule| %s\n", ruleCtx.GetText())
}
func (x *myListener) ExitEveryRule(ruleCtx antlr.ParserRuleContext) {
	ffmt.Printf("exit every rule| %s\n", ruleCtx.GetText())
}

func TestExampleListener(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want *parser.CypherParser
	}{
		{
			name: `MATCH`,
			args: args{query: `MATCH (:Person {name: "Alice"})-[:KNOWS]->(:Person {name: "Bob"}) RETURN count(*)`},
			want: nil,
		},
		{
			name: `CREATE`,
			args: args{query: `CREATE (n:Person {name: 'Andy', title: 'Developer'})`},
			want: nil,
		},
		{
			name: `MATCH`,
			args: args{query: `MATCH (p1:Person {name: "Alice", IsCool: $isCool})-[:KNOWS]->(p2:Person) WHERE p2.name <> "Tom" RETURN count(*)`},
			want: nil,
		},
		{
			name: `MATCH`,
			args: args{query: `MATCH (me)-[:KNOWS*1..2]-(remote_friend) WHERE me.name = 'Filipa' RETURN remote_friend.name`},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the input
			is := antlr.NewInputStream(tt.args.query)

			// Create the Lexer
			lexer := parser.NewCypherLexer(is)
			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

			// Create the Parser
			p := parser.NewCypherParser(stream)

			// Finally parse the expression (by walking the tree)
			// The "listener" is the action triggered when visiting each node
			listener := &myListener{}
			antlr.ParseTreeWalkerDefault.Walk(listener, p.OC_Cypher())
		})
	}
}
